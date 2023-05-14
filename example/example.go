package main

import (
	"fmt"

	"github.com/diebietse/gotp/v2"
	NorenApi "github.com/gauravjnigam/norenapigo"
)

func main() {

	// Create New Shoonya Broking Client
	ABClient := NorenApi.New("FA87226", "AlgoBaba@23", "aa4cff2b3742cc0eeeea60d51e311722")

	fmt.Println("Client :- ", ABClient)
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
	// fmt.Printf("Session : %v\n", session)
	// fmt.Printf("Token : %v\n", session.Susertoken)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	//Renew User Tokens using refresh token
	// session.Susertoken, err = ABClient.RenewAccessToken(session.Susertoken)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		// return
	}

	fmt.Println("User Session Tokens :- ", session.Susertoken)

	//Get User Profile
	session.UserProfile, err = ABClient.GetUserProfile()
	fmt.Println("User Profile :- ", session.UserProfile)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User Profile :- ", session.UserProfile)
	fmt.Println("User Session Object :- ", session)

	//Get Last Traded Price
	ltp, err := ABClient.GetLTP(NorenApi.LTPParams{Exchange: "NSE", Token: "3045"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Last Traded Price :- ", ltp)

	//Get Last Traded Price
	canldels, err := ABClient.GetTimePriceSeries(NorenApi.TSPriceParam{Exch: "NSE", Token: "Nifty Bank", St: "1683744945", Et: "1684018222", Intrv: "5"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(canldels)
}
