package main

import (
	"fmt"

	"github.com/diebietse/gotp/v2"
	NorenApi "github.com/gauravjnigam/norenapigo"
)

func main() {

	// Create New Shoonya Broking Client
	NorenClient := NorenApi.New("FA87226", "AlgoDada@23", "aa4cff2b3742cc0eeeea60d51e311722")

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

	// orders, err := NorenClient.GetOrderHistory(NorenApi.GetOrderParams{OrderID: "23071500006758"})

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Order hist resp : %v", orders)

	// //Get Last Traded Price
	//jData={"uid":"FA87226","actid":"FA87226","exch":"NSE","tsym":"HDFCBANK-EQ","qty":"1","prc":"1660.75","dscqty":"0","prd":"C","trantype":"B","prctyp":"LMT","ret":"DAY","ordersource":"WEB"}&jKey=eda03a4f0f1ad6937c9d5c208b40cef476c76ebcc3010409ef695049b994fd19
	// orderParam := NorenApi.OrderParams{
	// 	OrderSource:       "API",
	// 	UserId:            session.UID,
	// 	AccountId:         session.Actid,
	// 	TransactionType:   "B",
	// 	ProductType:       "C",
	// 	Exchange:          "NSE",
	// 	TradingSymbol:     "HDFCBANK-EQ",
	// 	Quantity:          "50",
	// 	PriceType:         "LMT",
	// 	Price:             "1000",
	// 	Retention:         "DAY",
	// 	Remarks:           "Test order",
	// 	DisclosedQuantity: "0",
	// }

	// ordResp, err := NorenClient.PlaceOrder(orderParam)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Order resp : %v", ordResp)

	// gttReqContext := NorenApi.GTTRequestContext{
	// 	Exchange:        "NFO",
	// 	TradingSymbol:   "BANKNIFTY03AUG23P46000",
	// 	TransactionType: "B",
	// 	AlertType:       "LTP_A_O",
	// 	AlertPriceAbove: 100,
	// 	AlertPriceBelow: 80,
	// 	PriceType:       "LMT",
	// 	Price:           0,
	// 	ProductType:     "M",
	// 	Quantity:        50,
	// 	Retention:       "DAY",
	// 	Discloseqty:     0,
	// }
	// gttResp, err := NorenClient.PlaceGTTOrder(gttReqContext)
	// if err != nil {
	// 	fmt.Printf("Error while placing GTT - %v\n", err)
	// }

	// fmt.Printf("GTT Response : %v", gttResp)

	gtts, err := NorenClient.GetPendingGTTOrder()
	if err != nil {
		fmt.Printf("Error - %v\n", err)
	}
	fmt.Printf("Pending GTTs - %v\n", gtts)

	gtts, err = NorenClient.CancelGTTOrder("23072900000311")
	if err != nil {
		fmt.Printf("Error - %v\n", err)
	}
	fmt.Printf("Cancelled GTTs - %v\n", gtts)

}
