package models

import "time"

// OHLC represents OHLC packets.
type OHLC struct {
	InstrumentToken uint32  `json:"-"`
	Open            float64 `json:"open"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Close           float64 `json:"close"`
}

// DepthItem represents a single market depth entry.
type DepthItem struct {
	Price    float64 `json:"price"`
	Quantity uint32  `json:"quantity"`
	Orders   uint32  `json:"orders"`
}

// Depth represents a group of buy/sell market depths.
type Depth struct {
	Buy  [5]DepthItem `json:"buy"`
	Sell [5]DepthItem `json:"sell"`
}

// Time is custom time format used in all responses
type Time struct {
	time.Time
}

type Tick struct {
	Timestamp     string `json:"ft"`
	Type          string `json:"t"`
	Exch          string `json:"e"`
	Token         string `json:"tk"`
	TradingSymbol string `json:"ts"`
	Pp            string `json:"pp"`
	Ls            string `json:"ls"`
	Ti            string `json:"ti"`
	LatestPrice   string `json:"lp"`
	Pc            string `json:"pc"`
	Open          string `json:"o"`
	High          string `json:"h"`
	Low           string `json:"l"`
	Close         string `json:"c"`
}

type OrderTick struct {
	Type       string `json:"t"`
	Norenordno string `json:"norenordno"`
	UID        string `json:"uid"`
	Actid      string `json:"actid"`
	Exch       string `json:"exch"`
	Tsym       string `json:"tsym"`
	Trantype   string `json:"trantype"`
	Qty        string `json:"qty"`
	Prc        string `json:"prc"`
	Pcode      string `json:"pcode"`
	Remarks    string `json:"remarks"`
	Status     string `json:"status"`
	Reporttype string `json:"reporttype"`
	Prctyp     string `json:"prctyp"`
	Ret        string `json:"ret"`
	Exchordid  string `json:"exchordid"`
	Dscqty     string `json:"dscqty"`
	Rejreason  string `json:"rejreason"`
}

// Tick represents a single packet in the market feed.
type Tick1 struct {
	Mode            string
	InstrumentToken uint32
	IsTradable      bool
	IsIndex         bool

	// Timestamp represents Exchange timestamp
	Timestamp          Time
	LastTradeTime      Time
	LastPrice          float64
	LastTradedQuantity uint32
	TotalBuyQuantity   uint32
	TotalSellQuantity  uint32
	VolumeTraded       uint32
	TotalBuy           uint32
	TotalSell          uint32
	AverageTradePrice  float64
	OI                 uint32
	OIDayHigh          uint32
	OIDayLow           uint32
	NetChange          float64

	OHLC  OHLC
	Depth Depth
}
