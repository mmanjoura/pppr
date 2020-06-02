package transaction

type Serializer interface {
	Decode(input []byte) (*Meta, error)
	Encode(input *Meta) ([]byte, error)
}
