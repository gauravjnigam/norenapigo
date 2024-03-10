package norenapigo

/*
func (ts *TestSuite) TestGetOrders(t *testing.T) {
	t.Parallel()
	gOrderParam := GetOrderParams{OrderID: "MOCK_ORD_ID"}
	orders, err := ts.TestConnect.GetOrderHistory(gOrderParam)
	if err != nil {
		t.Errorf("Error while fetching orders. %v", err)
	}

	ord := orders.(Order)
	if ord.OrderID == "" {
		t.Errorf("Error while fetching order id in orders. %v", err)
	}

}

func (ts *TestSuite) TestGetTrades(t *testing.T) {
	t.Parallel()
	trades, err := ts.TestConnect.GetTradeBook()
	if err != nil {
		t.Errorf("Error while fetching trades. %v", err)
	}
	for _, trade := range trades {
		if trade.OrderID == "" {
			t.Errorf("Error while fetching trade id in trades. %v", err)
		}
	}
}

func (ts *TestSuite) TestPlaceOrder(t *testing.T) {
	t.Parallel()
	params := OrderParams{"NORMAL", "SBIN-EQ", "3045", "BUY", "NSE", "LIMIT", "INTRADAY", "DAY", "19500", "0", "0", "1", ""}
	orderResponse, err := ts.TestConnect.PlaceOrder(params)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
	}
	if orderResponse.OrderId == "" {
		t.Errorf("No order id returned. Error %v", err)
	}
}

// func (ts *TestSuite) TestModifyOrder(t *testing.T) {
// 	t.Parallel()
// 	params := ModifyOrderParams{"NORMAL", "test", "LIMIT", "INTRADAY", "DAY", "19400", "1", "SBI-EQ", "3045", "NSE", "", ""}
// 	orderResponse, err := ts.TestConnect.ModifyOrder(params)
// 	if err != nil {
// 		t.Errorf("Error while updating order. %v", err)
// 	}
// 	if orderResponse.OrderId == "" {
// 		t.Errorf("No order id returned. Error %v", err)
// 	}
// }

func (ts *TestSuite) TestCancelOrder(t *testing.T) {
	t.Parallel()
	parentOrderID := "test"

	orderResponse, err := ts.TestConnect.CancelOrder("NORMAL", parentOrderID)
	if err != nil {
		t.Errorf("Error while cancellation of an order. %v", err)
	}
	if orderResponse.OrderId == "" {
		t.Errorf("No order id returned. Error %v", err)
	}
}

func (ts *TestSuite) TestGetPositions(t *testing.T) {
	t.Parallel()
	positions, err := ts.TestConnect.GetPositions()
	if err != nil {
		t.Errorf("Error while fetching positions. %v", err)
	}
	for _, position := range positions {
		if position.Exchange == "" {
			t.Errorf("Error while fetching exchange in positions. %v", err)
		}
	}
}

func (ts *TestSuite) TestConvertPosition(t *testing.T) {
	t.Parallel()
	params := ConvertPositionParams{"NSE", "SBIN-EQ", "DELIVERY", "MARGIN", "BUY", 1, "DAY"}
	err := ts.TestConnect.ConvertPosition(params)
	if err != nil {
		t.Errorf("Error while fetching positions. %v", err)
	}

}
*/
