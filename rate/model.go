package rate

// ExchangeRateTimedCube ...
type ExchangeRateTimedCube struct {
	CreatedDate string                     `json:"createddate,omitempty"`
	CreatedTime string                     `json:"createdtime,omitempty"`
	Rates       []ExchangeRateCurrencyCube `json:"rates,omitempty"`
}

// ExchangeRateCurrencyCube ...
type ExchangeRateCurrencyCube struct {
	Currency string  `json:"currency,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
}
