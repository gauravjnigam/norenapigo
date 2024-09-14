package upstock

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UpstoxClient struct {
	BaseURL string
}

// Define the structure for historical data response
type CandleData struct {
	TimeStamp string  `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type DailyTSData struct {
	Candles []CandleData
}

type HistoricalDataResponse struct {
	Status string `json:"status"`
	Data   struct {
		Candles [][]interface{} `json:"candles"`
	} `json:"data"`
}

func NewClient() *UpstoxClient {
	return &UpstoxClient{
		BaseURL: "https://api.upstox.com/v2/",
	}
}

// FetchHistoricalData fetches historical data for the given instrument
func (c *UpstoxClient) FetchHistoricalData(exchange, instrument, interval, startDate, endDate string) (*DailyTSData, error) {
	dailyTSDataResp := DailyTSData{}
	url := "https://api.upstox.com/v2/historical-candle"

	endPoint := fmt.Sprintf("%s/%s|%s/%s/%s/%s/", url, exchange, instrument, interval, endDate, startDate)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, endPoint, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	fmt.Printf("Req - %v\n", req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var historicalData HistoricalDataResponse
	err = json.NewDecoder(res.Body).Decode(&historicalData)
	if err != nil {
		return nil, err
	}
	var candles []CandleData
	for _, candle := range historicalData.Data.Candles {
		if len(candle) >= 6 {
			// Type assertions for each field from the interface{} slice
			timestamp, ok := candle[0].(string)
			if !ok {
				log.Println("error typecasting timestamp")
				continue
			}
			open, ok := candle[1].(float64)
			if !ok {
				log.Println("error typecasting open price")
				continue
			}
			high, ok := candle[2].(float64)
			if !ok {
				log.Println("error typecasting high price")
				continue
			}
			low, ok := candle[3].(float64)
			if !ok {
				log.Println("error typecasting low price")
				continue
			}
			closePrice, ok := candle[4].(float64)
			if !ok {
				log.Println("error typecasting close price")
				continue
			}
			volume, ok := candle[5].(float64)
			if !ok {
				log.Println("error typecasting volume")
				continue
			}

			// Append to the candles slice
			candles = append(candles, CandleData{
				TimeStamp: timestamp,
				Open:      open,
				High:      high,
				Low:       low,
				Close:     closePrice,
				Volume:    volume,
			})
		}
	}

	dailyTSDataResp.Candles = candles
	return &dailyTSDataResp, nil

}
