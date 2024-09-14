package main

import (
	"fmt"
	"log"

	"github.com/gauravjnigam/norenapigo/upstock"
)

func main() {
	client := upstock.NewClient()

	// exchange := "NSE_INDEX" // Set the exchange
	// instrument := "Nifty 50"

	exchange := "NSE_EQ" // Set the exchange
	instrument := "INE062A01020"

	interval := "day"
	startDate := "2024-08-26"
	endDate := "2024-09-11"

	// Fetch daily price data
	data, err := client.FetchHistoricalData(exchange, instrument, interval, startDate, endDate)
	if err != nil {
		log.Fatalf("error fetching data: %v", err)
	}

	// Print the structured data
	for _, c := range data.Candles {
		fmt.Printf("Timestamp: %s, Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f, Volume: %.2f\n",
			c.TimeStamp, c.Open, c.High, c.Low, c.Close, c.Volume)
	}

}
