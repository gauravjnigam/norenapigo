package norenapigo

import (
	"testing"
)

func (ts *TestSuite) TestGetHoldings(t *testing.T) {
	t.Parallel()
	holdings, err := ts.TestConnect.GetHoldings()
	if err != nil {
		t.Errorf("Error while fetching holdings. %v", err)
	}

	for _, holding := range holdings {
		if holding.Exchange == "" {
			t.Errorf("Error while fetching exchange in holdings. %v", err)
		}
	}

}
