package norenapigo

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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

type ExpiryDateResponse struct {
	ExpiryDate string `json:"expiryDate"`
}

type DailyTSParam struct {
	Symbol string `json:"sym"`
	St     string `json:"from"`
	Et     string `json:"to"`
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

// Get historial timePrice series
func (c *Client) GetDailyTimeSeries(exchange string, symbol string, startTime string, endTime string) (TSPResponse, error) {
	// fmt.Printf("Starttime: %s, EndTime: %s\n", startTime, endTime)
	// start, _ := GetDate(startTime)
	// end, _ := GetDate(endTime)
	end, start := GetTodayAndLastWeekEpoch()
	sym := fmt.Sprintf("%s:%s", exchange, symbol)
	fmt.Printf("Sym: %s, StartDate: %d, EndDate: %d\n", sym, start, end)
	tsPriceParam := DailyTSParam{Symbol: sym, St: fmt.Sprintf("%d", start), Et: fmt.Sprintf("%d", end)}
	var candle []Candle
	// fmt.Printf("Daily TSP Param: \n%v\n", tsPriceParam)
	params := structToMap(tsPriceParam, "json")
	fmt.Printf("Req Param : \n%v\n", params)
	params["uid"] = c.clientCode
	err := c.doEodEnvelope(http.MethodGet, URIEODTPSeries, params, nil, &candle, true)
	if err != nil {
		fmt.Errorf("Error while fetching EOD chart : %v", err.Error())
	}
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
	fmt.Printf("Search resp : %v", searchResp)
	var token string
	if searchResp.Stat == "Ok" && len(searchResp.Values) > 0 {
		token = searchResp.Values[0].Token
	}

	params["uid"] = c.clientCode
	params["exch"] = exchange
	params["token"] = token

	err1 := c.doEnvelope(http.MethodPost, URILTP, params, nil, &ltp, true)
	fmt.Printf("LTP resp : %v", ltp)
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

func (c *Client) GetExpiry(basesearch string, instrumentCode string, optionOrFuture string, expiryType string, expiryInstance int) (ExpiryDateResponse, error) {

	var expiryResp ExpiryDateResponse

	searchString := c.getSearchStringForExpiry(basesearch, instrumentCode, optionOrFuture, expiryType)

	searchResp, err := c.Searchscrip("NFO", searchString)

	if err != nil {
		fmt.Errorf("Error while fetching expiry date - %v", err)
		return expiryResp, err
	}

	if searchResp.Stat == "Ok" && len(searchResp.Values) > 0 {
		for _, val := range searchResp.Values {
			fmt.Printf("%s\n", val.Tsym)
		}
	}

	return expiryResp, nil

}

func (c *Client) getSearchStringForExpiry(baseSearch string, instrumentCode string, optionOrFuture string, expiryType string) string {
	ltp, err := c.GetLTP(LTPParams{Exchange: "NSE", Token: instrumentCode})

	if err != nil {
		fmt.Errorf("Error while fetching LTP to calculate expiry - %v", err)
		return "--INVALID--"
	}

	// atm price for instument i.e. Nifty, banknifty, finnifty
	ltpFlt, _ := strconv.ParseFloat(ltp.C, 64)
	atmPrice := getATMStrike(ltpFlt, ltp.Tsym)
	// fmt.Printf("%s LTP : %v", ltp.Tsym, atmPrice)
	atmPriceStr := fmt.Sprintf("%d", atmPrice)
	t := time.Now()
	monthYear := fmt.Sprintf("%s%d", t.Month().String()[:3], t.Year()%1e2)
	// fmt.Printf("Current Month and year : %s\n", monthYear)

	if optionOrFuture == "OPTION" {
		baseSearch = baseSearch + " " + monthYear + " C" + " " + atmPriceStr
	}
	// fmt.Printf("Search string for expiry :\n" + baseSearch)
	return baseSearch
}

func (c *Client) GetExpiryDate(expiryType string, expiryInstance int) (ExpiryDateResponse, error) {
	var expiryDate time.Time
	var expDateResp ExpiryDateResponse
	dayVal := time.Now()
	// fmt.Printf("\nExpiry request for %s - %d, for day: %s\n", expiryType, expiryInstance, dayVal.Format("2006-01-02"))

	if expiryType == "WEEKLY" {
		// fmt.Printf("Weekly\n")
		dayVal = dayVal.Add(time.Duration(7*expiryInstance) * 24 * time.Hour)
		// fmt.Printf("dayVal - %v\n", dayVal)
		wexp := weeklyExpiry(dayVal)
		expiryDate = wexp
	}

	if expiryType == "MONTHLY" {
		dayVal = dayVal.Add(time.Duration(31*expiryInstance) * 24 * time.Hour)
		mexp := monthlyExpiry(dayVal)
		expiryDate = mexp
	}

	if !expiryDate.IsZero() {
		expDateStr := expiryDate.Format("02Jan06")
		// fmt.Printf("Expiry date = %s\n", expDateStr)
		expDateResp.ExpiryDate = expDateStr
	}

	return expDateResp, nil
}

func getATMStrike(ltp float64, instrumentSymbol string) int64 {
	if instrumentSymbol == "Nifty Bank" || instrumentSymbol == "Bank Nifty" {
		rem := int64(ltp) % 100
		return int64(ltp) + (100 - rem)
	} else if instrumentSymbol == "NIFTY" || instrumentSymbol == "Nifty 50" {
		return int64(ltp) + (50 - int64(ltp)%50)
	} else {
		rem := int64(ltp) % 100
		return int64(ltp) + (100 - rem)
	}
}
