package websocket

import (
	"bytes"
	"compress/zlib"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SocketClient struct {
	Conn                *websocket.Conn
	url                 url.URL
	callbacks           callbacks
	autoReconnect       bool
	reconnectMaxRetries int
	reconnectMaxDelay   time.Duration
	connectTimeout      time.Duration
	reconnectAttempt    int
	scrips              string
	feedToken           string
	clientCode          string
}

// callbacks represents callbacks available in ticker.
type callbacks struct {
	onMessage     func([]map[string]interface{})
	onNoReconnect func(int)
	onReconnect   func(int, time.Duration)
	onConnect     func()
	onClose       func(int, string)
	onError       func(error)
}

const (
	// Auto reconnect defaults
	// Default maximum number of reconnect attempts
	defaultReconnectMaxAttempts = 3
	// Auto reconnect min delay. Reconnect delay can't be less than this.
	reconnectMinDelay time.Duration = 5000 * time.Millisecond
	// Default auto reconnect delay to be used for auto reconnection.
	defaultReconnectMaxDelay time.Duration = 60000 * time.Millisecond
	// Connect timeout for initial server handshake.
	defaultConnectTimeout time.Duration = 7000 * time.Millisecond
	// Interval in which the connection check is performed periodically.
	connectionCheckInterval time.Duration = 10000 * time.Millisecond
)

var (
	// Default ticker url.
	tickerURL = url.URL{Scheme: "wss", Host: "api.shoonya.com", Path: "/NorenWSTP/"}
)

// New creates a new ticker instance.
func New(clientCode string, feedToken string, scrips string) *SocketClient {
	sc := &SocketClient{
		clientCode:          clientCode,
		feedToken:           feedToken,
		url:                 tickerURL,
		autoReconnect:       true,
		reconnectMaxDelay:   defaultReconnectMaxDelay,
		reconnectMaxRetries: defaultReconnectMaxAttempts,
		connectTimeout:      defaultConnectTimeout,
		scrips:              scrips,
	}

	// fmt.Printf("Socket client - %v\n", sc)

	return sc
}

// SetRootURL sets ticker root url.
func (s *SocketClient) SetRootURL(u url.URL) {
	s.url = u
}

// SetAccessToken set access token.
func (s *SocketClient) SetFeedToken(feedToken string) {
	s.feedToken = feedToken
}

// SetConnectTimeout sets default timeout for initial connect handshake
func (s *SocketClient) SetConnectTimeout(val time.Duration) {
	s.connectTimeout = val
}

// SetAutoReconnect enable/disable auto reconnect.
func (s *SocketClient) SetAutoReconnect(val bool) {
	s.autoReconnect = val
}

// SetReconnectMaxDelay sets maximum auto reconnect delay.
func (s *SocketClient) SetReconnectMaxDelay(val time.Duration) error {
	if val > reconnectMinDelay {
		return fmt.Errorf("ReconnectMaxDelay can't be less than %fms", reconnectMinDelay.Seconds()*1000)
	}

	s.reconnectMaxDelay = val
	return nil
}

// SetReconnectMaxRetries sets maximum reconnect attempts.
func (s *SocketClient) SetReconnectMaxRetries(val int) {
	s.reconnectMaxRetries = val
}

// OnConnect callback.
func (s *SocketClient) OnConnect(f func()) {
	fmt.Println("Connecting...")
	s.callbacks.onConnect = f

}

// OnError callback.
func (s *SocketClient) OnError(f func(err error)) {
	fmt.Println("Errored out...")
	s.callbacks.onError = f
}

// OnClose callback.
func (s *SocketClient) OnClose(f func(code int, reason string)) {
	fmt.Println("Closed...")
	s.callbacks.onClose = f
}

// OnMessage callback.
func (s *SocketClient) OnMessage(f func(message []map[string]interface{})) {
	fmt.Println("OnMessage...")
	s.callbacks.onMessage = f

}

// OnReconnect callback.
func (s *SocketClient) OnReconnect(f func(attempt int, delay time.Duration)) {
	fmt.Println("Reconnecting")
	s.callbacks.onReconnect = f
}

// OnNoReconnect callback.
func (s *SocketClient) OnNoReconnect(f func(attempt int)) {
	fmt.Println("Not connected")
	s.callbacks.onNoReconnect = f
}

// Serve starts the connection to ticker server. Since its blocking its recommended to use it in go routine.
func (s *SocketClient) Serve() {
	fmt.Println("\n\n\n serving...")

	for {
		fmt.Println("--> ")
		// If reconnect attempt exceeds max then close the loop
		if s.reconnectAttempt > s.reconnectMaxRetries {
			s.triggerNoReconnect(s.reconnectAttempt)
			return
		}
		// If its a reconnect then wait exponentially based on reconnect attempt
		if s.reconnectAttempt > 0 {
			nextDelay := time.Duration(math.Pow(2, float64(s.reconnectAttempt))) * time.Second
			if nextDelay > s.reconnectMaxDelay {
				nextDelay = s.reconnectMaxDelay
			}

			s.triggerReconnect(s.reconnectAttempt, nextDelay)

			// Close the previous connection if exists
			if s.Conn != nil {
				s.Conn.Close()
			}
		}
		// create a dialer
		d := websocket.DefaultDialer
		d.HandshakeTimeout = s.connectTimeout
		d.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		fmt.Println("--> " + s.url.String())
		conn, _, err := d.Dial(s.url.String(), nil)
		// fmt.Printf("Resp : %v", resp)
		if err != nil {
			s.triggerError(err)
			// If auto reconnect is enabled then try reconneting else return error
			if s.autoReconnect {
				s.reconnectAttempt++
				continue
			}
			return
		}
		// fmt.Printf("Tick : %v\n", websocket.)
		params := map[string]interface{}{}
		params["t"] = "c"
		params["k"] = "NSE|26009"
		params["uid"] = s.clientCode
		params["actid"] = s.clientCode
		params["susertoken"] = s.feedToken
		params["source"] = "API"
		jsonPayload, err := json.Marshal(params)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, jsonPayload)
		if err != nil {
			fmt.Printf("Error aa gaya : %v\n", err)
			s.triggerError(err)
			return
		}
		fmt.Println("Error ni aaya")
		_, message, err := conn.ReadMessage()
		if err != nil {
			s.triggerError(err)
			return
		}

		sDec, _ := base64.StdEncoding.DecodeString(string(message))
		val, err := readSegment(sDec)
		var result []map[string]interface{}
		err = json.Unmarshal(val, &result)
		fmt.Printf("Tick : %v\n", string(message))
		fmt.Printf("val : %v\n", val)
		fmt.Printf("Result : %v\n", result)

		if err != nil {
			s.triggerError(err)
			return
		}
		fmt.Printf("Message : %v\n", result)
		if len(result) == 0 {
			s.triggerError(fmt.Errorf("Invalid Message"))
			return
		}

		if _, ok := result[0]["ak"]; !ok {
			s.triggerError(fmt.Errorf("Invalid Message"))
			return
		}

		if val, ok := result[0]["ak"]; ok {
			if val == "nk" {
				s.triggerError(fmt.Errorf("Invalid feed token or client code"))
				return
			}
		}

		// Close the connection when its done.
		defer s.Conn.Close()

		// Assign the current connection to the instance.
		s.Conn = conn

		// Trigger connect callback.
		s.triggerConnect()

		// Resubscribe to stored tokens
		if s.reconnectAttempt > 0 {
			_ = s.Resubscribe()
		}

		// Reset auto reconnect vars
		s.reconnectAttempt = 0

		// Set on close handler
		s.Conn.SetCloseHandler(s.handleClose)

		var wg sync.WaitGroup
		Restart := make(chan bool, 1)
		// Receive ticker data in a go routine.
		wg.Add(1)
		go s.readMessage(&wg, Restart)

		// Run watcher to check last ping time and reconnect if required
		if s.autoReconnect {
			wg.Add(1)
			go s.checkConnection(&wg, Restart)
		}

		// Wait for go routines to finish before doing next reconnect
		wg.Wait()
	}
}

func (s *SocketClient) handleClose(code int, reason string) error {
	s.triggerClose(code, reason)
	return nil
}

// Trigger callback methods
func (s *SocketClient) triggerError(err error) {
	if s.callbacks.onError != nil {
		s.callbacks.onError(err)
	}
}

func (s *SocketClient) triggerClose(code int, reason string) {
	if s.callbacks.onClose != nil {
		s.callbacks.onClose(code, reason)
	}
}

func (s *SocketClient) triggerConnect() {
	if s.callbacks.onConnect != nil {
		s.callbacks.onConnect()
	}
}

func (s *SocketClient) triggerReconnect(attempt int, delay time.Duration) {
	if s.callbacks.onReconnect != nil {
		s.callbacks.onReconnect(attempt, delay)
	}
}

func (s *SocketClient) triggerNoReconnect(attempt int) {
	if s.callbacks.onNoReconnect != nil {
		s.callbacks.onNoReconnect(attempt)
	}
}

func (s *SocketClient) triggerMessage(message []map[string]interface{}) {
	if s.callbacks.onMessage != nil {
		s.callbacks.onMessage(message)
	}
}

// Periodically check for last ping time and initiate reconnect if applicable.
func (s *SocketClient) checkConnection(wg *sync.WaitGroup, Restart chan bool) {
	defer wg.Done()
	switch {
	case <-Restart:
		return
	}
}

// readMessage reads the data in a loop.
func (s *SocketClient) readMessage(wg *sync.WaitGroup, Restart chan bool) {
	defer wg.Done()
	for {
		fmt.Println("Msg aaya ")
		_, msg, err := s.Conn.ReadMessage()
		if err != nil {
			s.triggerError(fmt.Errorf("Error reading data: %v", err))
			Restart <- true
			return
		}
		fmt.Println("Msg : " + string(msg))
		sDec, _ := base64.StdEncoding.DecodeString(string(msg))
		val, err := readSegment(sDec)
		if err != nil {
			s.triggerError(err)
			return
		}

		var finalMessage []map[string]interface{}
		err = json.Unmarshal(val, &finalMessage)
		if err != nil {
			s.triggerError(err)
			return
		}

		if len(finalMessage) == 0 {
			continue
		}

		if val, ok := finalMessage[0]["ak"]; ok {
			if val == "nk" {
				s.triggerError(fmt.Errorf("Invalid feed token or client code"))
			}
			continue
		}

		// Trigger message.
		s.triggerMessage(finalMessage)

	}
}

// Close tries to close the connection gracefully. If the server doesn't close it
func (s *SocketClient) Close() error {
	return s.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

// Subscribe subscribes tick for the given list of tokens.
func (s *SocketClient) Subscribe() error {
	fmt.Println("Subscribing... ")
	err := s.Conn.WriteMessage(websocket.TextMessage, []byte(`{"task":"mw","channel":"`+s.scrips+`","token":"`+s.feedToken+`","user": "`+s.clientCode+`","acctid":"`+s.clientCode+`"}`))
	if err != nil {
		s.triggerError(err)
		return err
	}

	return nil
}

func (s *SocketClient) Resubscribe() error {
	err := s.Subscribe()
	return err
}

func readSegment(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)
	z, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	p, err := ioutil.ReadAll(z)
	if err != nil {
		return nil, err
	}
	return p, nil
}
