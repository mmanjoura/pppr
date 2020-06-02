package json

import (
	"encoding/json"

	"github.com/mmanjoura/pppr/logging"
	"github.com/pkg/errors"
)

// Logging ...
type Message struct{}

// Decode ...
func (r *Message) Decode(input []byte) (*logging.LogMessage, error) {
	lgn := &logging.LogMessage{}
	if err := json.Unmarshal(input, lgn); err != nil {
		return nil, errors.Wrap(err, "serializer.Logging.Decode")
	}
	return lgn, nil
}

// Encode ...
func (r *Message) Encode(input *logging.LogMessage) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Logging.Encode")
	}
	return rawMsg, nil
}
