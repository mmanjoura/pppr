package json

import (
	"encoding/json"

	"github.com/mmanjoura/pppr/payment"
	"github.com/pkg/errors"
)

// Payment ...
type Payment struct{}

// Decode ...
func (r *Payment) Decode(input []byte) (*payment.RunParams, error) {
	trx := &payment.RunParams{}
	if err := json.Unmarshal(input, trx); err != nil {
		return nil, errors.Wrap(err, "serializer.Payment.Decode")
	}
	return trx, nil
}

// Encode ...
func (r *Payment) Encode(input *payment.RunParams) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Payment.Encode")
	}
	return rawMsg, nil
}
