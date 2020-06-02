package payment

//Serializer ...
type Serializer interface {
	Decode(input []byte) (*RunParams, error)
	Encode(input *RunParams) ([]byte, error)
}
