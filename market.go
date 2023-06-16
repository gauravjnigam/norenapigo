package norenapigo

import (
	"net/http"
)

// LTPResponse represents LTP API Response.
type LTPResponse struct {
	Exch  string `json:"exch"`
	Tsym  string `json:"tsym"`
	Token string `json:"token"`
	Lp    string `json:"lp"`
	C     string `json:"c"`
	H     string `json:"h"`
	L     string `json:"l"`
	Ap    string `json:"ap"`
	O     string `json:"o"`
}

// LTPParams represents parameters for getting LTP.
type LTPParams struct {
	Exchange string `json:"exch"`
	Token    string `json:"token"`
}

type Candle struct {
	Stat    string `json:"stat"`
	Time    string `json:"time"`
	Ssboe   string `json:"ssboe"`
	Into    string `json:"into"`
	Inth    string `json:"inth"`
	Intl    string `json:"intl"`
	Intc    string `json:"intc"`
	Intvwap string `json:"intvwap"`
	Intv    string `json:"intv"`
	Intoi   string `json:"intoi"`
	V       string `json:"v"`
	Oi      string `json:"oi"`
}

type TSPResponse struct {
	Candles []Candle
}

type TSPriceParam struct {
	Exch  string `json:"exch"`
	Token string `json:"token"`
	St    string `json:"st"`
	Et    string `json:"et"`
	Intrv string `json:"intrv"`
}

// GetLTP gets Last Traded Price.
func (c *Client) GetLTP(ltpParams LTPParams) (LTPResponse, error) {
	var ltp LTPResponse
	params := structToMap(ltpParams, "json")
	params["uid"] = c.clientCode
	err := c.doEnvelope(http.MethodPost, URILTP, params, nil, &ltp, true)
	return ltp, err
}

// Get historial timePrice series
func (c *Client) GetTimePriceSeries(exchange string, token string, startTime string, endTime string, interval string) (TSPResponse, error) {
	start := GetTime(startTime)
	end := GetTime(endTime)
	tsPriceParam := TSPriceParam{Exch: exchange, Token: token, St: start, Et: end, Intrv: interval}
	var candle []Candle
	// fmt.Printf("TSP Param: \n%v\n", tsPriceParam)
	params := structToMap(tsPriceParam, "json")
	// fmt.Printf("Req Param : \n%v\n", params)
	params["uid"] = c.clientCode
	err := c.doEnvelope(http.MethodPost, URITPSeries, params, nil, &candle, true)
	tsResponse := TSPResponse{Candles: candle}
	return tsResponse, err
}
