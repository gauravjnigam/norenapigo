package main

import (
	"fmt"
	"time"

	"github.com/diebietse/gotp/v2"
	"github.com/norenapigo/v2"
	NorenApi "github.com/norenapigo/v2"
	models "github.com/norenapigo/v2/model"
	"github.com/norenapigo/v2/websocket"
)

var socketClient *websocket.SocketClient

// Triggered when tick is recevived
func onTick(tick models.Tick) {
	fmt.Println("Tick: ", tick)
}

// Triggered when order update is received
func onOrderUpdate(order norenapigo.Order) {
	fmt.Printf("Order: %s", order.OrderID)
}

// Triggered when any error is raised
func onError(err error) {
	fmt.Println("Error aayo re: ", err)
}

// Triggered when websocket connection is closed
func onClose(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnect() {
	fmt.Println("Connected")
	// fmt.Printf("SocketClient : %v\n", socketClient)
	err := socketClient.Subscribe()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when a message is received
func onMessage(message []byte) {
	fmt.Printf("Message Received :- %v\n", string(message))
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnect(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnect(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d\n", attempt)
}

// var wg sync.WaitGroup

func main() {

	// Create New Angel Broking Client
	ABClient := NorenApi.New("FA87226", "AlgoBaba@23", "aa4cff2b3742cc0eeeea60d51e311722")

	// fmt.Println("Client :- ", ABClient)
	clientTotpSecret := "U6CFCE65M63MLV655H25D2327HU36YYJ"
	secret, err := gotp.DecodeBase32(clientTotpSecret)
	if err != nil {
		panic(err)
	}
	otp, err := gotp.NewTOTP(secret)
	if err != nil {
		panic(err)
	}
	currentOTP, err := otp.Now()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("current one-time password is: %v\n", currentOTP)

	// User Login and Generate User Session
	session, err := ABClient.GenerateSession(currentOTP)
	// fmt.Printf("WS : Session : %v\n", session)
	fmt.Printf("WS : Token : %v\n", session.Susertoken)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
	fmt.Printf("Starting websocket connection \n [%v, %v]\n", session.UID, session.Susertoken)

	// time.Sleep(1 * time.Second)
	// New Websocket Client
	socketClient = websocket.New(session.UID, session.Susertoken, "NSE|26009#NSE|22")

	// Assign callbacks
	socketClient.OnError(onError)
	socketClient.OnClose(onClose)
	socketClient.OnMessage(onMessage)
	socketClient.OnConnect(onConnect)
	socketClient.OnReconnect(onReconnect)
	socketClient.OnNoReconnect(onNoReconnect)
	socketClient.OnTick(onTick)
	socketClient.OnOrderUpdate(onOrderUpdate)

	// Start Consuming Data
	// wg.Add(1)
	// fmt.Printf("SocketClient1 : %v\n", socketClient)
	socketClient.Serve()
	// wg.Wait()
	// fmt.Println("Done")

}
