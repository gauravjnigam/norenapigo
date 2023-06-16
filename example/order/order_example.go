package main

import (
	"fmt"

	"github.com/diebietse/gotp/v2"
	NorenApi "github.com/gauravjnigam/norenapigo"
)

func main() {

	// Create New Shoonya Broking Client
	NorenClient := NorenApi.New("FA87226", "AlgoBaba@23", "aa4cff2b3742cc0eeeea60d51e311722")

	fmt.Println("Client :- ", NorenClient)
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
	session, err := NorenClient.GenerateSession(currentOTP)
	// fmt.Printf("Session : %v\n", session)
	// fmt.Printf("Token : %v\n", session.Susertoken)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	//Renew User Tokens using refresh token
	// session.Susertoken, err = NorenClient.RenewAccessToken(session.Susertoken)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		// return
	}

	fmt.Println("User Session Tokens :- ", session.Susertoken)

	//Get User Profile
	session.UserProfile, err = NorenClient.GetUserProfile()
	fmt.Println("User Profile :- ", session.UserProfile)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User Profile :- ", session.UserProfile)
	fmt.Println("User Session Object :- ", session)

	// //Get Last Traded Price
	// orderParam := NorenApi.OrderParams{
	// 	OrderSource: "API",
	// }
	// _, err := NorenClient.PlaceOrder(orderParam)

}