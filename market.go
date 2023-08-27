package norenapigo

import (
	"fmt"
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

// LTPParams represents parameters for getting LTP.
type SecurityInfoParams struct {
	Exchange string `json:"exch"`
	Token    string `json:"token"`
}

type SecurityInfoResponse struct {
	Exchange      string `json:"exch"`
	TradingSymbol string `json:"tsym"`
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

type SearchResponse struct {
	Stat   string `json:"stat"`
	Values []struct {
		Exch  string `json:"exch"`
		Token string `json:"token"`
		Tsym  string `json:"tsym"`
	} `json:"values"`
	ErrorMessage string `json:"emsg"`
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

// GetLTP gets Last Traded Price.
func (c *Client) GetSecurityInfo(secInfoParam SecurityInfoParams) (SecurityInfoResponse, error) {
	var siResponse SecurityInfoResponse
	params := structToMap(secInfoParam, "json")
	params["uid"] = c.clientCode
	err := c.doEnvelope(http.MethodPost, URISecurityInfo, params, nil, &siResponse, true)
	return siResponse, err
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

// GetTradingSymbol(self, exchange, searchText)
func (c *Client) GetTradingSymbol(exchange string, searchText string) (string, error) {
	searchResp, err := c.Searchscrip(exchange, searchText)
	if err != nil {
		fmt.Printf("Error while searching - %v", err)
	}
	var tsym string
	if searchResp.Stat == "Ok" && len(searchResp.Values) > 0 {
		tsym = searchResp.Values[0].Tsym
	} else {
		return "", err
	}

	return tsym, nil
}

// GetLatestPrice(tradingSymbol, exchange):
func (c *Client) GetLatestPrice(tradingSymbol string, exchange string) (LTPResponse, error) {
	var ltp LTPResponse
	params := map[string]interface{}{}
	searchResp, err := c.Searchscrip(exchange, tradingSymbol)
	if err != nil {
		fmt.Printf("Error while searching - %v", err)
	}
	var token string
	if searchResp.Stat == "Ok" && len(searchResp.Values) > 0 {
		token = searchResp.Values[0].Token
	}

	params["uid"] = c.clientCode
	params["exch"] = exchange
	params["token"] = token

	err1 := c.doEnvelope(http.MethodPost, URILTP, params, nil, &ltp, true)

	return ltp, err1
}

func (c *Client) Searchscrip(exchange string, searchscrip string) (SearchResponse, error) {

	var searchResponse SearchResponse
	params := map[string]interface{}{}
	params["uid"] = c.clientCode
	params["exch"] = exchange
	params["stext"] = searchscrip

	err := c.doEnvelope(http.MethodPost, URISearchScript, params, nil, &searchResponse, true)

	return searchResponse, err
}

// def GetExpriyDate(self, expiryType, expiryInstance):
// func (c *Client) GetExpriyDate(expiryType string, expiryInstance string) (string, error) {

// }
