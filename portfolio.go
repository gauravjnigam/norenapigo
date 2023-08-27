package norenapigo

import (
	"net/http"
)

// Holding is an individual holdings response.
type Holding struct {
	Tradingsymbol      string `json:"tradingsymbol"`
	Exchange           string `json:"exchange"`
	ISIN               string `json:"isin"`
	T1Quantity         string `json:"t1quantity"`
	RealisedQuantity   string `json:"realisedquantity"`
	Quantity           string `json:"quantity"`
	AuthorisedQuantity string `json:"authorisedquantity"`
	ProfitAndLoss      string `json:"profitandloss"`
	Product            string `json:"product"`
	CollateralQuantity string `json:"collateralquantity"`
	CollateralType     string `json:"collateraltype"`
	Haircut            string `json:"haircut"`
}

// Holdings is a list of holdings
type Holdings []Holding

// GetHoldings gets a list of holdings.
func (c *Client) GetHoldings() (Holdings, error) {
	var holdings Holdings

	params := map[string]interface{}{}
	params["uid"] = c.clientCode
	params["actid"] = c.clientCode
	params["prd"] = "C"

	err := c.doEnvelope(http.MethodPost, URIGetHoldings, params, nil, &holdings, true)
	return holdings, err
}
