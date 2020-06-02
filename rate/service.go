package rate

// Service ...
type Service interface {
	Save(rates ExchangeRateTimedCube, collection string) error
}
