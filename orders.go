package norenapigo

import (
	"fmt"
	"net/http"
	"strconv"
)

type GetOrderParams struct {
	OrderID string `json:"norenordno"`
}

// Order represents a individual order response.
type Order struct {
	Variety                 string `json:"variety"`
	OrderType               string `json:"ordertype"`
	ProductType             string `json:"producttype"`
	Duration                string `json:"duration"`
	Price                   string `json:"price"`
	TriggerPrice            string `json:"triggerprice"`
	Quantity                string `json:"quantity"`
	DisclosedQuantity       string `json:"disclosedquantity"`
	SquareOff               string `json:"squareoff"`
	StopLoss                string `json:"stoploss"`
	TrailingStopLoss        string `json:"trailingstoploss"`
	TrailingSymbol          string `json:"trailingsymbol"`
	TransactionType         string `json:"transactiontype"`
	Exchange                string `json:"exchange"`
	SymbolToken             string `json:"symboltoken"`
	InstrumentType          string `json:"instrumenttype"`
	StrikePrice             string `json:"strikeprice"`
	OptionType              string `json:"optiontype"`
	ExpiryDate              string `json:"expirydate"`
	LotSize                 string `json:"lotsize"`
	CancelSize              string `json:"cancelsize"`
	AveragePrice            string `json:"averageprice"`
	FilledShares            string `json:"filledshares"`
	UnfilledShares          string `json:"unfilledshares"`
	OrderID                 string `json:"orderid"`
	Text                    string `json:"text"`
	Status                  string `json:"status"`
	OrderStatus             string `json:"orderstatus"`
	UpdateTime              string `json:"updatetime"`
	ExchangeTime            string `json:"exchtime"`
	ExchangeOrderUpdateTime string `json:"exchorderupdatetime"`
	FillID                  string `json:"fillid"`
	FillTime                string `json:"filltime"`
}

// Orders is a list of orders.
type Orders []Order

// OrderParams represents parameters for placing an order.
type OrderParams struct {
	OrderSource       string `json:"ordersource"`
	UserId            string `json:"uid"`
	AccountId         string `json:"actid"`
	TransactionType   string `json:"trantype"`
	ProductType       string `json:"prd"`
	Exchange          string `json:"exch"`
	TradingSymbol     string `json:"tsym "`
	Quantity          string `json:"qty"`
	PriceType         string `json:"prctyp"`
	Price             string `json:"price"`
	Retention         string `json:"ret"`
	Remarks           string `json:"remarks"`
	DisclosedQuantity string `json:"dscqty"`
}

// OrderParams represents parameters for modifying an order.
type ModifyOrderParams struct {
	OrderSource     string `json:"ordersource"`
	UserId          string `json:"uid"`
	AccountId       string `json:"actid"`
	TransactionType string `json:"trantype"`
	ProductType     string `json:"prd"`
	Exchange        string `json:"exch"`
	TradingSymbol   string `json:"tsym "`
	Quantity        string `json:"qty"`
	PriceType       string `json:"prctyp"`
	Price           string `json:"prc"`
	Retention       string `json:"ret"`
	Remarks         string `json:"remarks"`
}

// OrderResponse represents the order place success response.
type OrderResponse struct {
	OrderStatus  string `json:"stat"`
	OrderId      string `json:"norenordno"`
	RequestTime  string `json:"request_time"`
	ErrorMessage string `json:"emsg"`
}

type GTTOrderResponse struct {
	Stat        string `json:"stat"`
	AlertID     string `json:"Al_id"`
	RequestTime string `json:"request_time"`
	Emsg        string `json:"emsg"`
}

type GTTRequestContext struct {
	Exchange        string
	TradingSymbol   string
	TransactionType string
	AlertType       string
	AlertPriceAbove float64
	AlertPriceBelow float64
	PriceType       string
	Price           float64
	ProductType     string
	Quantity        int
	Retention       string
	Validity        string
	OrderRemark     string
	Discloseqty     int
}

//	product_type= gCtx.product_type, quantity= gCtx.quantity, price_type=gCtx.price_type, price=gCtx.price, remarks= gCtx.order_remark, retention=gCtx.retention, validity=gCtx.validity, discloseqty=gCtx.discloseqty
//
// GTTOrderParams represents parameters for placing a GTT order.
type GTTOrderParams struct {
	AlertType       string
	TradingSymbol   string
	Exchange        string
	AlertPrice      float64
	AlertPriceAbove float64
	AlertPriceBelow float64
	TransactionType string
	ProductType     string
	Quantity        int
	PriceType       string
	Price           float64
	Remarks         string
	Retention       string
	Validity        string
	Discloseqty     int
}

// Trade represents an individual trade response.
type Trade struct {
	Exchange        string `json:"exchange"`
	ProductType     string `json:"producttype"`
	TradingSymbol   string `json:"tradingsymbol"`
	InstrumentType  string `json:"instrumenttype"`
	SymbolGroup     string `json:"symbolgroup"`
	StrikePrice     string `json:"strikeprice"`
	OptionType      string `json:"optiontype"`
	ExpiryDate      string `json:"expirydate"`
	MarketLot       string `json:"marketlot"`
	Precision       string `json:"precision"`
	Multiplier      string `json:"multiplier"`
	TradeValue      string `json:"tradevalue"`
	TransactionType string `json:"transactiontype"`
	FillPrice       string `json:"fillprice"`
	FillSize        string `json:"fillsize"`
	OrderID         string `json:"orderid"`
	FillID          string `json:"fillid"`
	FillTime        string `json:"filltime"`
}

// Trades is a list of trades.
type Trades []Trade

// Position represents an individual position response.
type Position struct {
	Exchange              string `json:"exchange"`
	SymbolToken           string `json:"symboltoken"`
	ProductType           string `json:"producttype"`
	Tradingsymbol         string `json:"tradingsymbol"`
	SymbolName            string `json:"symbolname"`
	InstrumentType        string `json:"instrumenttype"`
	PriceDen              string `json:"priceden"`
	PriceNum              string `json:"pricenum"`
	GenDen                string `json:"genden"`
	GenNum                string `json:"gennum"`
	Precision             string `json:"precision"`
	Multiplier            string `json:"multiplier"`
	BoardLotSize          string `json:"boardlotsize"`
	BuyQuantity           string `json:"buyquantity"`
	SellQuantity          string `json:"sellquantity"`
	BuyAmount             string `json:"buyamount"`
	SellAmount            string `json:"sellamount"`
	SymbolGroup           string `json:"symbolgroup"`
	StrikePrice           string `json:"strikeprice"`
	OptionType            string `json:"optiontype"`
	ExpiryDate            string `json:"expirydate"`
	LotSize               string `json:"lotsize"`
	CfBuyQty              string `json:"cfbuyqty"`
	CfSellQty             string `json:"cfsellqty"`
	CfBuyAmount           string `json:"cfbuyamount"`
	CfSellAmount          string `json:"cfsellamount"`
	BuyAveragePrice       string `json:"buyavgprice"`
	SellAveragePrice      string `json:"sellavgprice"`
	AverageNetPrice       string `json:"avgnetprice"`
	NetValue              string `json:"netvalue"`
	NetQty                string `json:"netqty"`
	TotalBuyValue         string `json:"totalbuyvalue"`
	TotalSellValue        string `json:"totalsellvalue"`
	CfBuyAveragePrice     string `json:"cfbuyavgprice"`
	CfSellAveragePrice    string `json:"cfsellavgprice"`
	TotalBuyAveragePrice  string `json:"totalbuyavgprice"`
	TotalSellAveragePrice string `json:"totalsellavgprice"`
	NetPrice              string `json:"netprice"`
}

// Positions represents a list of net and day positions.
type Positions []Position

// ConvertPositionParams represents the input params for a position conversion.
type ConvertPositionParams struct {
	Exchange        string `url:"exchange"`
	TradingSymbol   string `url:"tradingsymbol"`
	OldProductType  string `url:"oldproducttype"`
	NewProductType  string `url:"newproducttype"`
	TransactionType string `url:"transactiontype"`
	Quantity        int    `url:"quantity"`
	Type            string `json:"type"`
}

// GetOrderBook gets user orders.
func (c *Client) GetOrderBook() (interface{}, error) {
	var orders interface{}
	params := map[string]interface{}{}
	params["uid"] = c.clientCode

	err := c.doEnvelope(http.MethodPost, URIGetOrderBook, params, nil, &orders, true)
	return orders, err
}

// GetOrderBook gets user orders.
func (c *Client) GetOrderHistory(getOrderParams GetOrderParams) (interface{}, error) {
	var orders interface{}
	params := structToMap(getOrderParams, "json")
	params["uid"] = c.clientCode
	params["actid"] = c.clientCode

	err := c.doEnvelope(http.MethodPost, URIGetOrderBook, params, nil, &orders, true)
	return orders, err
}

// PlaceOrder places an order.
func (c *Client) PlaceOrder(orderParams OrderParams) (OrderResponse, error) {
	var orderResponse OrderResponse
	// params := structToMap(orderParams, "json")
	params := map[string]interface{}{}
	params["uid"] = c.clientCode

	params["ordersource"] = "API"
	params["uid"] = c.clientCode
	params["actid"] = orderParams.AccountId
	params["trantype"] = orderParams.TransactionType
	params["prd"] = orderParams.ProductType
	params["exch"] = orderParams.Exchange
	params["tsym"] = orderParams.TradingSymbol
	params["qty"] = orderParams.Quantity
	params["dscqty"] = orderParams.DisclosedQuantity
	params["prctyp"] = orderParams.PriceType
	params["prc"] = orderParams.Price
	params["trgprc"] = ""
	params["ret"] = orderParams.Retention
	params["remarks"] = orderParams.Remarks
	params["amo"] = "NO"

	fmt.Printf("Param : %v", params)
	err := c.doEnvelope(http.MethodPost, URIPlaceOrder, params, nil, &orderResponse, true)

	return orderResponse, err
}

// ModifyOrder for modifying an order.
func (c *Client) ModifyOrder(orderParams ModifyOrderParams) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        map[string]interface{}
		err           error
	)

	params["uid"] = c.clientCode

	params["ordersource"] = "API"
	params["uid"] = c.clientCode
	params["actid"] = orderParams.AccountId
	params["trantype"] = orderParams.TransactionType
	params["prd"] = orderParams.ProductType
	params["exch"] = orderParams.Exchange
	params["tsym"] = orderParams.TradingSymbol
	params["qty"] = orderParams.Quantity
	params["dscqty"] = "0"
	params["prctyp"] = orderParams.PriceType
	params["prc"] = orderParams.Price
	params["trgprc"] = ""
	params["ret"] = orderParams.Retention
	params["remarks"] = orderParams.Remarks
	params["amo"] = "NO"

	err = c.doEnvelope(http.MethodPost, URIModifyOrder, params, nil, &orderResponse, true)
	return orderResponse, err
}

// CancelOrder for cancellation of an order.
func (c *Client) CancelOrder(variety string, orderid string) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		err           error
	)

	params := make(map[string]interface{})
	params["variety"] = variety
	params["orderid"] = orderid

	err = c.doEnvelope(http.MethodPost, URICancelOrder, params, nil, &orderResponse, true)
	return orderResponse, err
}

// GetPositions gets user positions.
func (c *Client) GetPositions() (Positions, error) {
	var positions Positions
	err := c.doEnvelope(http.MethodGet, URIGetPositions, nil, nil, &positions, true)
	return positions, err
}

// GetTradeBook gets user trades.
func (c *Client) GetTradeBook() (Trades, error) {
	var trades Trades
	err := c.doEnvelope(http.MethodGet, URIGetTradeBook, nil, nil, &trades, true)
	return trades, err
}

// GTT Apis
// PlaceGTTOrder
func (c *Client) PlaceGTTOrder(gttCtx GTTOrderParams) (interface{}, error) {
	var gttOrderResponse interface{}

	if gttCtx.AlertType == "LTP_A_O" || gttCtx.AlertType == "LTP_B_O" {
		alertPrice := 0.0
		if gttCtx.AlertType == "LTP_A_O" {
			alertPrice = gttCtx.AlertPriceAbove
		}

		if gttCtx.AlertType == "LTP_B_O" {
			alertPrice = gttCtx.AlertPriceBelow
		}

		if gttCtx.PriceType == "LMT" {
			c.PlaceGTT_LMT_order(GTTOrderParams{TradingSymbol: gttCtx.TradingSymbol, Exchange: gttCtx.Exchange, AlertType: gttCtx.AlertType,
				AlertPrice: alertPrice, TransactionType: gttCtx.TransactionType,
				ProductType: gttCtx.ProductType,
				Quantity:    gttCtx.Quantity,
				PriceType:   gttCtx.PriceType,
				Price:       gttCtx.Price,
				Remarks:     gttCtx.Remarks,
				Retention:   "DAY",
				Validity:    "GTT",
				Discloseqty: gttCtx.Discloseqty})

		} else if gttCtx.PriceType == "MKT" {
			c.PlaceGTT_MKT_order(GTTOrderParams{TradingSymbol: gttCtx.TradingSymbol, Exchange: gttCtx.Exchange, AlertType: gttCtx.AlertType,
				AlertPrice: alertPrice, TransactionType: gttCtx.TransactionType,
				ProductType: gttCtx.ProductType,
				Quantity:    gttCtx.Quantity,
				PriceType:   gttCtx.PriceType,
				Price:       gttCtx.Price,
				Remarks:     gttCtx.Remarks,
				Retention:   "DAY",
				Validity:    "GTT",
				Discloseqty: gttCtx.Discloseqty})
		} else {

		}

	} else if gttCtx.AlertType == "LMT_BOS_O" {
		if gttCtx.PriceType == "LMT" {
			c.PlaceGTT_OCO_LMT_order(GTTOrderParams{TradingSymbol: gttCtx.TradingSymbol, Exchange: gttCtx.Exchange, AlertType: gttCtx.AlertType,
				AlertPriceAbove: gttCtx.AlertPriceAbove, AlertPriceBelow: gttCtx.AlertPriceBelow, TransactionType: gttCtx.TransactionType,
				ProductType: gttCtx.ProductType,
				Quantity:    gttCtx.Quantity,
				PriceType:   gttCtx.PriceType,
				Price:       gttCtx.Price,
				Remarks:     gttCtx.Remarks,
				Retention:   "DAY",
				Validity:    "GTT",
				Discloseqty: gttCtx.Discloseqty})
		} else if gttCtx.PriceType == "MKT" {
			c.PlaceGTT_OCO_MKT_order(GTTOrderParams{TradingSymbol: gttCtx.TradingSymbol, Exchange: gttCtx.Exchange, AlertType: gttCtx.AlertType,
				AlertPriceAbove: gttCtx.AlertPriceAbove, AlertPriceBelow: gttCtx.AlertPriceBelow, TransactionType: gttCtx.TransactionType,
				ProductType: gttCtx.ProductType,
				Quantity:    gttCtx.Quantity,
				PriceType:   gttCtx.PriceType,
				Price:       gttCtx.Price,
				Remarks:     gttCtx.Remarks,
				Retention:   "DAY",
				Validity:    "GTT",
				Discloseqty: gttCtx.Discloseqty})
		} else {

		}
	}

	return gttOrderResponse, nil
}

func (c *Client) PlaceGTT_LMT_order(gttOrdParam GTTOrderParams) (GTTOrderResponse, error) {
	fmt.Printf("Placing GTT LMT Order for req : %v\n", gttOrdParam)
	var gttOrdResponse GTTOrderResponse

	params := map[string]interface{}{}
	params["ordersource"] = "API"
	params["uid"] = c.clientCode
	params["actid"] = c.clientCode
	params["tsym"] = gttOrdParam.TradingSymbol
	params["exch"] = gttOrdParam.Exchange
	params["ai_t"] = gttOrdParam.AlertType
	params["validity"] = gttOrdParam.Validity
	var ap string = strconv.FormatFloat(gttOrdParam.AlertPrice, 'E', -1, 32)
	params["d"] = ap
	params["remarks"] = gttOrdParam.Remarks
	params["trantype"] = gttOrdParam.TransactionType
	params["prctyp"] = gttOrdParam.PriceType
	params["prd"] = gttOrdParam.ProductType
	params["ret"] = gttOrdParam.Retention
	var qty_str string = strconv.Itoa(gttOrdParam.Quantity)
	params["qty"] = qty_str
	var price_str string = strconv.FormatFloat(gttOrdParam.Price, 'E', -1, 32)
	params["prc"] = price_str
	var dscqty_str string = strconv.Itoa(gttOrdParam.Discloseqty)
	params["dscqty"] = dscqty_str

	err := c.doEnvelope(http.MethodPost, URIPlaceGTTOrder, params, nil, &gttOrdResponse, true)
	return gttOrdResponse, err

}
func (c *Client) PlaceGTT_MKT_order(gttOrdParam GTTOrderParams) (GTTOrderResponse, error) {
	var (
		gttOrdResponse GTTOrderResponse
		params         map[string]interface{}
		err            error
	)

	params["uid"] = c.clientCode

	err = c.doEnvelope(http.MethodPost, URIPlaceGTTOrder, params, nil, &gttOrdResponse, true)
	return gttOrdResponse, err
}
func (c *Client) PlaceGTT_OCO_LMT_order(gttOrdParam GTTOrderParams) (GTTOrderResponse, error) {
	var (
		gttOrdResponse GTTOrderResponse
		params         map[string]interface{}
		err            error
	)

	params["uid"] = c.clientCode

	err = c.doEnvelope(http.MethodPost, URIPlaceGTTOrder, params, nil, &gttOrdResponse, true)
	return gttOrdResponse, err
}
func (c *Client) PlaceGTT_OCO_MKT_order(gttOrdParam GTTOrderParams) (GTTOrderResponse, error) {
	var (
		gttOrdResponse GTTOrderResponse
		params         map[string]interface{}
		err            error
	)

	params["uid"] = c.clientCode

	err = c.doEnvelope(http.MethodPost, URIPlaceGTTOrder, params, nil, &gttOrdResponse, true)
	return gttOrdResponse, err
}

// ModifyGTTOrder
func (c *Client) ModifyGTTOrder(gttCtx GTTRequestContext) (interface{}, error) {
	return nil, nil
}

// CancelGTTOrder
func (c *Client) CancelGTTOrder(alertId string) (interface{}, error) {
	var cancelGTTOrderResp interface{}
	params := make(map[string]interface{})
	params["uid"] = c.clientCode
	params["al_id"] = alertId

	err := c.doEnvelope(http.MethodPost, URICancelGTTOrder, params, nil, &cancelGTTOrderResp, true)

	if err != nil {
		fmt.Printf("Error while cancelling GTT for alert id - %v\n", alertId)
		return nil, err
	}
	return cancelGTTOrderResp, nil
}

// GetPendingGTTOrder
func (c *Client) GetPendingGTTOrder() (interface{}, error) {
	var pendingGTTOrders interface{}
	params := make(map[string]interface{})
	params["uid"] = c.clientCode

	err := c.doEnvelope(http.MethodPost, URIGetPendingGTTOrder, params, nil, &pendingGTTOrders, true)

	if err != nil {
		fmt.Printf("Error while getting pending GTTs - %v\n", err)
		return nil, err
	}
	return pendingGTTOrders, nil
}

// GetEnabledGTTs
func (c *Client) GetEnabledGTTs() (interface{}, error) {
	var enabledGTTResp interface{}
	params := make(map[string]interface{})
	params["uid"] = c.clientCode

	err := c.doEnvelope(http.MethodPost, URIGetEnabledGTTs, params, nil, &enabledGTTResp, true)

	if err != nil {
		fmt.Printf("Error while getting enabled GTT response - %v\n", err)
		return nil, err
	}
	return enabledGTTResp, nil
}
