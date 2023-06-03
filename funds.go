package norenapigo

// RMS represents API response.
type RMS struct {
	Net                    string `json:"net"`
	AvailableCash          string `json:"availablecash"`
	AvailableIntraDayPayIn string `json:"availableintradaypayin"`
	AvailableLimitMargin   string `json:"availablelimitmargin"`
	Collateral             string `json:"collateral"`
	M2MUnrealized          string `json:"m2munrealized"`
	M2MRealized            string `json:"m2mrealized"`
	UtilisedDebits         string `json:"utiliseddebits"`
	UtilisedSpan           string `json:"utilisedspan"`
	UtilisedOptionPremium  string `json:"utilisedoptionpremium"`
	UtilisedHoldingSales   string `json:"utilisedholdingsales"`
	UtilisedExposure       string `json:"utilisedexposure"`
	UtilisedTurnover       string `json:"utilisedturnover"`
	UtilisedPayout         string `json:"utilisedpayout"`
}
