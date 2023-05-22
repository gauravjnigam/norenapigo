package websocket

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/norenapigo/v2"
	models "github.com/norenapigo/v2/model"
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
	cancel              context.CancelFunc
}

// callbacks represents callbacks available in ticker.
type callbacks struct {
	onMessage     func([]map[string]interface{})
	onNoReconnect func(int)
	onReconnect   func(int, time.Duration)
	onConnect     func()
	onClose       func(int, string)
	onError       func(error)
	onTick        func(models.Tick)
	onOrderUpdate func(norenapigo.Order)
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

	// fmt.Printf("Socket client url - %v\n", tickerURL)

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
	// fmt.Println("Connecting...")
	s.callbacks.onConnect = f

}

// OnError callback.
func (s *SocketClient) OnError(f func(err error)) {
	// fmt.Println("Errored out...")
	s.callbacks.onError = f
}

// OnTick callback.
func (s *SocketClient) OnTick(f func(tick models.Tick)) {
	// fmt.Println("Errored out...")
	s.callbacks.onTick = f
}

// OnTick callback.
func (s *SocketClient) OnOrderUpdate(f func(order norenapigo.Order)) {
	// fmt.Println("Errored out...")
	s.callbacks.onOrderUpdate = f
}

// OnClose callback.
func (s *SocketClient) OnClose(f func(code int, reason string)) {
	// fmt.Println("Closed...")
	s.callbacks.onClose = f
}

// OnMessage callback.
func (s *SocketClient) OnMessage(f func(message []map[string]interface{})) {
	fmt.Println("OnMessage...")
	s.callbacks.onMessage = f

}

// OnReconnect callback.
func (s *SocketClient) OnReconnect(f func(attempt int, delay time.Duration)) {
	// fmt.Println("Reconnecting")
	s.callbacks.onReconnect = f
}

// OnNoReconnect callback.
func (s *SocketClient) OnNoReconnect(f func(attempt int)) {
	// fmt.Println("Not connected")
	s.callbacks.onNoReconnect = f
}

// Serve starts the connection to ticker server. Since its blocking its
// recommended to use it in a go routine.
func (s *SocketClient) Serve() {
	s.ServeWithContext(context.Background())
}

// ServeWithContext starts the connection to ticker server and additionally
// accepts a context. Since its blocking its recommended to use it in a go
// routine.
func (s *SocketClient) ServeWithContext(ctx context.Context) {
	fmt.Println("\nServing...")
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	for {
		// fmt.Println("--> ")
		select {
		case <-ctx.Done():
			return
		default:
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

			q := s.url.Query()
			q.Set("t", "c")
			q.Set("actid", s.clientCode)
			q.Set("uid", s.clientCode)
			q.Set("access_token", s.feedToken)
			q.Set("source", "API")

			s.url.RawQuery = q.Encode()
			// create a dialer
			d := websocket.DefaultDialer
			d.HandshakeTimeout = s.connectTimeout
			// d.TLSClientConfig = &tls.Config{
			// 	InsecureSkipVerify: true,
			// }
			fmt.Printf("URL --> %v\n", s.url)

			conn, _, err := d.Dial(s.url.String(), nil)
			// fmt.Printf("\nResp : %v\n", resp)
			// fmt.Printf("\nConn : %v\n", conn)
			if err != nil {
				s.triggerError(err)
				// If auto reconnect is enabled then try reconneting else return error
				if s.autoReconnect {
					s.reconnectAttempt++
					continue
				}
				return
			}

			fmt.Printf("Dialed ... \n")
			// Close the connection when its done.
			defer s.Conn.Close()

			// Assign the current connection to the instance.
			s.Conn = conn

			// Trigger connect callback.
			s.triggerConnect()

			// Resubscribe to stored tokens
			if s.reconnectAttempt > 0 {
				fmt.Printf("Reconnecting ... \n")
				_ = s.Resubscribe()
			}

			// Reset auto reconnect vars
			s.reconnectAttempt = 0

			// Set on close handler
			s.Conn.SetCloseHandler(s.handleClose)
			fmt.Printf("Waiting for message ... \n")
			var wg sync.WaitGroup
			Restart := make(chan bool, 1)
			// Receive ticker data in a go routine.
			wg.Add(1)
			go s.readMessage(ctx, &wg, Restart)

			// Run watcher to check last ping time and reconnect if required
			if s.autoReconnect {
				wg.Add(1)
				go s.checkConnection(&wg, Restart)
			}

			// Wait for go routines to finish before doing next reconnect
			wg.Wait()
		}
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
func (s *SocketClient) readMessage(ctx context.Context, wg *sync.WaitGroup, Restart chan bool) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// fmt.Println("Msg ekde ")
			mType, msg, err := s.Conn.ReadMessage()
			fmt.Println("Msg read bhayo re ")
			if err != nil {
				s.triggerError(fmt.Errorf("Error reading data: %v", err))
				Restart <- true
				return
			}
			fmt.Println("Msg : " + string(msg))
			if mType == websocket.BinaryMessage {
				fmt.Printf("MsgType : %v\n", mType)
			} else {
				fmt.Printf("MsgType : %v\n", mType)
			}

			sDec, _ := base64.StdEncoding.DecodeString(string(msg))
			fmt.Printf("Msg : %v\n", string(msg))
			fmt.Printf("sDec : %v\n", sDec)
			val, err := readSegment(sDec)
			if err != nil {
				s.triggerError(err)
				return
			}
			fmt.Println("3")
			var finalMessage []map[string]interface{}
			err = json.Unmarshal(val, &finalMessage)
			if err != nil {
				s.triggerError(err)
				return
			}
			fmt.Println("4")
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
}

// func (t *SocketClient) processTextMessage(inp []byte) {
// 	var msg message
// 	if err := json.Unmarshal(inp, &msg); err != nil {
// 		// May be error should be triggered
// 		return
// 	}

// 	if msg.Type == messageError {
// 		// Trigger text error
// 		t.triggerError(fmt.Errorf(msg.Data.(string)))
// 	} else if msg.Type == messageOrder {
// 		// Parse order update data
// 		order := struct {
// 			Data norenapigo.Order `json:"data"`
// 		}{}

// 		if err := json.Unmarshal(inp, &order); err != nil {
// 			// May be error should be triggered
// 			return
// 		}

// 		t.triggerOrderUpdate(order.Data)
// 	}
// }

// Close tries to close the connection gracefully. If the server doesn't close it
func (s *SocketClient) Close() error {
	return s.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

type tickerInput struct {
	Type string `json:"t"`
	Val  string `json:"k"`
}

type orderInput struct {
	Type string      `json:"t"`
	Val  interface{} `json:"accid"`
}

// Subscribe subscribes tick for the given list of tokens.
func (s *SocketClient) Subscribe() error {
	fmt.Println("Subscribing... ")

	out, err := json.Marshal(orderInput{
		Type: "o",
		Val:  s.clientCode, //"NSE|20009",
	})
	if err != nil {
		return err
	}

	fmt.Printf("Payload : %v\n", string(out))
	// return s.Conn.WriteJSON(out)
	s.Conn.NextWriter(1)
	return s.Conn.WriteMessage(websocket.TextMessage, out)
	// if err != nil {
	// 	fmt.Printf("Error aa gaya : %v\n", err)
	// 	s.triggerError(err)
	// 	// return
	// }

	// return nil
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
