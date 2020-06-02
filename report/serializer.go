package report

// Serializer ...
type Serializer interface {
	EncodeGenerate(input Report) ([]byte, error)
	DecodeGenerate(input []byte) ([]Merchant, error)
}
