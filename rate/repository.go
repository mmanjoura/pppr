package rate

// Repository ...
type Repository interface {
	Save(rates ExchangeRateTimedCube, collection string) error
}
