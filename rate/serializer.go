package rate

// Serializer ...
type Serializer interface {
	Decode(input []byte) (ExchangeRateTimedCube, error)
	Encode(input ExchangeRateTimedCube) ([]byte, error)
}
