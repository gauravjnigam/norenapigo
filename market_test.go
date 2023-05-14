package norenapigo

import (
	"testing"
)

func (ts *TestSuite) TestGetLTP(t *testing.T) {
	t.Parallel()
	params := LTPParams{
		Exchange:      "NSE",
		TradingSymbol: "SBIN-EQ",
		SymbolToken:   "3045",
	}
	ltp, err := ts.TestConnect.GetLTP(params)
	if err != nil {
		t.Errorf("Error while fetching LTP. %v", err)
	}

	if ltp.Exchange == "" {
		t.Errorf("Error while exchange in LTP. %v", err)
	}

}
