package json

import (
	"encoding/json"

	"github.com/mmanjoura/pppr/report"
	"github.com/pkg/errors"
)

type Report struct{}

func (r *Report) DecodeGenerate(input []byte) ([]report.Merchant, error) {
	m := []report.Merchant{}
	if err := json.Unmarshal(input, m); err != nil {
		return nil, errors.Wrap(err, "serializer.Report.DecodeGeerate")
	}
	return m, nil
}

// EncodeGetMrchants ...
func (r *Report) EncodeGenerate(input report.Report) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Report.DecodeGeerate")
	}
	return rawMsg, nil
}
