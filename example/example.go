package main

import (
	"fmt"

	"github.com/diebietse/gotp/v2"
	NorenApi "github.com/gauravjnigam/norenapigo"
)

func main() {

	// Create New Shoonya Broking Client
	NorenClient := NorenApi.New("FA87226", "AlgoKaka@24", "aa4cff2b3742cc0eeeea60d51e311722")

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

	//Get Last Traded Price
	ltp, err := NorenClient.GetLTP(NorenApi.LTPParams{Exchange: "NSE", Token: "3045"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Last Traded Price :- ", ltp)
	// start := "09-06-2023 09:15:00"
	// end := "09-06-2023 15:10:00"

	// //Get Last Traded Price
	// fmt.Println("Fetching timeseries data - ")
	// // canldels, err := NorenClient.GetTimePriceSeries("NSE", "26000", start, end, "60")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(canldels)

	orders, err := NorenClient.GetSecurityInfo(NorenApi.SecurityInfoParams{Exchange: "NSE", Token: "3045"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("SecurityInfo resp : %v", orders)

	holdings, err1 := NorenClient.GetHoldings()
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Printf("Holdings resp : %v", holdings)

	searchScriptRes, err1 := NorenClient.Searchscrip("NFO", "NIFTY 07MAR24 22000")
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Printf("Searched text resp : %v\n", searchScriptRes)

	tsym := "nifty07MAR24P22000"
	ltpResp, err1 := NorenClient.GetLatestPrice(tsym, "NFO")
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Printf("Latest price for %s : %v", tsym, ltpResp.Token)

}
