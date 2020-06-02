package json

import (
	"encoding/json"

	"github.com/mmanjoura/pppr/acquirer"
	"github.com/pkg/errors"
)

// Payment ...
type Acquirer struct{}

func (r *Acquirer) DecodeGetAcquirerPayments(input []byte) ([]acquirer.Payment, error) {
	m := []acquirer.Payment{}
	if err := json.Unmarshal(input, m); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetAcquirerPayments")
	}
	return m, nil
}
func (r *Acquirer) DecodeGetMerchantPayments(input []byte) ([]acquirer.Payment, error) {
	m := []acquirer.Payment{}
	if err := json.Unmarshal(input, m); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetMerchantPayments")
	}
	return m, nil
}

func (r *Acquirer) DecodeGetReports(input []byte) ([]acquirer.Report, error) {
	rep := []acquirer.Report{}
	if err := json.Unmarshal(input, rep); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetReports")
	}
	return rep, nil
}

func (r *Acquirer) DecodeGetTransactions(input []byte) ([]acquirer.Transaction, error) {
	t := []acquirer.Transaction{}
	if err := json.Unmarshal(input, t); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetTransactions")
	}
	return t, nil
}

func (r *Acquirer) DecodeGetExchangeRates(input []byte) (acquirer.ExchangeRate, error) {
	exr := acquirer.ExchangeRate{}
	if err := json.Unmarshal(input, exr); err != nil {
		return acquirer.ExchangeRate{}, errors.Wrap(err, "serializer.Acquirer.DecodeGetExchangeRates")
	}
	return exr, nil
}

func (r *Acquirer) DecodeGetLogMessages(input []byte) ([]acquirer.Message, error) {
	lm := []acquirer.Message{}
	if err := json.Unmarshal(input, lm); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetLogMessages")
	}
	return lm, nil
}

func (r *Acquirer) DecodeGetProcessStates(input []byte) ([]acquirer.ProcessState, error) {
	ps := []acquirer.ProcessState{}
	if err := json.Unmarshal(input, ps); err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.DecodeGetProcessStates")
	}
	return ps, nil
}

func (r *Acquirer) DecodePutProcessState(input []byte) error {
	ps := []acquirer.ProcessState{}
	if err := json.Unmarshal(input, ps); err != nil {
		return errors.Wrap(err, "serializer.Acquirer.DecodePutProcessState")
	}
	return nil
}

func (r *Acquirer) EncodeGetAcquirerPayments(input []acquirer.Payment) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetAcquirerPayments")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetMerchantPayments(input []acquirer.Payment) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetMerchantPayments")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetReports(input []acquirer.Report) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetReports")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetTransactions(input []acquirer.Transaction) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetTransactions")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetExchangeRates(input acquirer.ExchangeRate) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetExchangeRates")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetLogMessages(input []acquirer.Message) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetLogMessages")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodeGetProcessStates(input []acquirer.ProcessState) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Acquirer.EncodeGetProcessStates")
	}
	return rawMsg, nil
}

func (r *Acquirer) EncodePutProcessState(input acquirer.ProcessState) error {

	return nil
}

// func (r *Acquirer) EndcodePostAcquirer(input Acquirer) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodeGetAcquirer(input Acquirer) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodePutAcquirer(input Acquirer) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodePatchAcquirer(input Acquirer) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodeDeleteAcquirer(input Acquirer) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }

// func (r *Acquirer) DecodePostAcquirer(input []byte) (acquirer.Acquirer, error) {
// 	a := acquirer.Acquirer{}
// 	return a, nil
// }
// func (r *Acquirer) DecodeGetAcquirer(input []byte) (acquirer.Acquirer, error) {
// 	a := acquirer.Acquirer{}
// 	return a, nil
// }
// func (r *Acquirer) DecodePutAcquirer(input []byte) (acquirer.Acquirer, error) {
// 	a := acquirer.Acquirer{}
// 	return a, nil
// }
// func (r *Acquirer) DecodePatchAcquirer(input []byte) (acquirer.Acquirer, error) {
// 	a := acquirer.Acquirer{}
// 	return a, nil
// }
// func (r *Acquirer) DecodeDeleteAcquirer(input []byte) (acquirer.Acquirer, error) {
// 	a := acquirer.Acquirer{}
// 	return a, nil
// }

// func (r *Acquirer) EndcodeGetMerchant(input acquirer.Merchant) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodePostMerchant(input acquirer.Merchant) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodeDeleteMerchant(input acquirer.Merchant) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }

// func (r *Acquirer) DecodeGetMerchant(input []byte) (acquirer.Merchant, error) {
// 	a := acquirer.Merchant{}
// 	return a, nil
// }
// func (r *Acquirer) DecodePostMerchant(input []byte) (acquirer.Merchant, error) {
// 	a := acquirer.Merchant{}
// 	return a, nil
// }
// func (r *Acquirer) DecodeDeleteMerchant(input []byte) (acquirer.Merchant, error) {
// 	a := acquirer.Merchant{}
// 	return a, nil
// }

// func (r *Acquirer) EndcodeGetTransactions(input []acquirer.Transaction) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodeGetTransaction(input acquirer.Transaction) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodePostTransaction(input acquirer.Transaction) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }
// func (r *Acquirer) EndcodeDeleteTransaction(input acquirer.Transaction) ([]byte, error) {
// 	rawMsg, err := json.Marshal(input)
// 	return rawMsg, err
// }

// func (r *Acquirer) DecodeGetTransactions(input []byte) ([]acquirer.Transaction, error) {
// 	a := []acquirer.Transaction{}
// 	return a, nil
// }
// func (r *Acquirer) DecodeGetTransaction(input []byte) (acquirer.Transaction, error) {
// 	a := acquirer.Transaction{}
// 	return a, nil
// }
// func (r *Acquirer) DecodePostTransaction(input []byte) (acquirer.Transaction, error) {
// 	a := acquirer.Transaction{}
// 	return a, nil
// }
// func (r *Acquirer) DecodeDeleteTransaction(input []byte) (acquirer.Transaction, error) {
// 	a := acquirer.Transaction{}
// 	return a, nil
// }
