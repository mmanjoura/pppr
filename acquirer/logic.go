package acquirer

import (
	"errors"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrTransactionNotFound  = errors.New("Transaction Not Found")
	ErrReportNotFound       = errors.New("Report Not Found")
	ErrExchangeRateNotFound = errors.New("ExchangeRate Not Found")
	ErrLogMessageNotFound   = errors.New("LogMessage Not Found")
	ErrProcessStateNotFound = errors.New("ProcessType Not Found")
	// ErrAcquirerNotFound    = errors.New("Acquirer Not Found")
	// ErrMerchantNotFound    = errors.New("Merchant Not Found")

	ErrTransactionInvalid  = errors.New("Transaction Invalid")
	ErrExchangeRateInvalid = errors.New("ExchangeRate Invalid")
	ErrLogMessageInvalid   = errors.New("LogMessage Invalid")
	ErrProcessStateInvalid = errors.New("ProcessState Invalid")
	// ErrTransactionIdInvalid = errors.New("Transaction ID Invalid")
	// ErrMerchantInvalid      = errors.New("Merchant  Invalid")
	// ErrAcquirerInvalid      = errors.New("Acquirer  Invalid")
	ErrAcquirerIdInvalid     = errors.New("Acquirer ID Invalid")
	ErrProcessStateIdInvalid = errors.New("Process State ID Invalid")
	ErrMerchantIdInvalid     = errors.New("Merchant ID Invalid")
	// ErrMerchantIdInvalid    = errors.New("Merchant ID Invalid")

	ErrInconsistentIDs = errors.New("inconsistent IDs")
	// ErrAlreadyExists   = errors.New("already exists")
	// ErrNotFound        = errors.New("not found")
)

type acquirerService struct {
	acquirerRepo Repository
}

// NewAcquirerService ...
func NewAcquirerService(acquirerRepo Repository) Service {
	return &acquirerService{
		acquirerRepo,
	}
}

func (r *acquirerService) GetAcquirerPayments(acquirerID string) ([]Payment, error) {
	if err := validate.Validate(acquirerID); err != nil {
		return nil, errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.GetMerchants")
	}

	payments, err := r.acquirerRepo.GetAcquirerPayments(acquirerID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *acquirerService) GetMerchantPayments(MID string) ([]Payment, error) {
	if err := validate.Validate(MID); err != nil {
		return nil, errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.GetMerchantPayments")
	}

	payments, err := r.acquirerRepo.GetMerchantPayments(MID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *acquirerService) GetReports(acquirerID string) ([]Report, error) {
	if err := validate.Validate(acquirerID); err != nil {
		return nil, errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.GetReports")
	}

	reports, err := r.acquirerRepo.GetReports(acquirerID)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *acquirerService) GetTransactions(merchantID string) ([]Transaction, error) {
	if err := validate.Validate(merchantID); err != nil {
		return []Transaction{}, errs.Wrap(ErrTransactionInvalid, "service.Acquirer.GetTransactions")
	}
	transactions, err := r.acquirerRepo.GetTransactions(merchantID)
	if err != nil {
		return []Transaction{}, ErrTransactionNotFound
	}
	return transactions, nil
}

func (r *acquirerService) GetExchangeRates(date string) (ExchangeRate, error) {
	if err := validate.Validate(date); err != nil {
		return ExchangeRate{}, errs.Wrap(ErrExchangeRateInvalid, "service.Acquirer.GetExchangeRates")
	}
	exchangeRates, err := r.acquirerRepo.GetExchangeRates(date)
	if err != nil {
		return ExchangeRate{}, ErrExchangeRateNotFound
	}
	return exchangeRates, nil
}

func (r *acquirerService) GetLogMessages(date string) ([]Message, error) {
	if err := validate.Validate(date); err != nil {
		return []Message{}, errs.Wrap(ErrLogMessageInvalid, "service.Acquirer.GetLogMessages")
	}
	logMessages, err := r.acquirerRepo.GetLogMessages(date)
	if err != nil {
		return []Message{}, ErrLogMessageNotFound
	}
	return logMessages, nil
}

func (r *acquirerService) GetProcessStates(approved bool) ([]ProcessState, error) {
	if err := validate.Validate(approved); err != nil {
		return []ProcessState{}, errs.Wrap(ErrProcessStateInvalid, "service.Acquirer.GetProcessStates")
	}
	processStates, err := r.acquirerRepo.GetProcessStates(approved)
	if err != nil {
		return []ProcessState{}, ErrProcessStateNotFound
	}
	return processStates, nil
}

func (r *acquirerService) PutProcessState(processID string, ps ProcessState) error {
	if processID != ps.ID {
		return ErrInconsistentIDs
	}
	if err := validate.Validate(processID); err != nil {
		return errs.Wrap(ErrProcessStateIdInvalid, "service.Acquirer.PutProcessState")
	}
	if err := validate.Validate(ps); err != nil {
		return errs.Wrap(ErrProcessStateInvalid, "service.Acquirer.PutProcessState")
	}
	err := r.acquirerRepo.PutProcessState(processID, ps) // PUT = create or update
	if err != nil {
		return err
	}
	return nil
}

// func (r *acquirerService) PostAcquirer(a Acquirer) error {
// 	if err := validate.Validate(a); err != nil {
// 		return errs.Wrap(ErrAcquirerInvalid, "service.Acquirer.PostAcquirer")
// 	}
// 	if err := r.acquirerRepo.PostAcquirer(a); err != nil {
// 		return ErrAlreadyExists // POST = create, don't overwrite
// 	}

// 	return nil
// }

// func (r *acquirerService) GetAcquirer(id string) (Acquirer, error) {
// 	if err := validate.Validate(id); err != nil {
// 		return Acquirer{}, errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.GetAcquirer")
// 	}
// 	acquirer, err := r.acquirerRepo.GetAcquirer(id)
// 	if err != nil {
// 		return Acquirer{}, ErrNotFound
// 	}
// 	return acquirer, nil
// }

// func (r *acquirerService) PutAcquirer(id string, a Acquirer) error {
// 	if id != a.ID {
// 		return ErrInconsistentIDs
// 	}
// 	if err := validate.Validate(id); err != nil {
// 		return errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.PutAcquirer")
// 	}
// 	if err := validate.Validate(a); err != nil {
// 		return errs.Wrap(ErrAcquirerInvalid, "service.Acquirer.PutAcquirer")
// 	}
// 	err := r.acquirerRepo.PutAcquirer(id, a) // PUT = create or update
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *acquirerService) PatchAcquirer(id string, a Acquirer) error {
// 	if a.ID != "" && id != a.ID {
// 		return ErrInconsistentIDs
// 	}

// 	if err := validate.Validate(id); err != nil {
// 		return errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.PatchAcquirer")
// 	}
// 	if err := validate.Validate(a); err != nil {
// 		return errs.Wrap(ErrAcquirerInvalid, "service.Acquirer.PatchAcquirer")
// 	}

// 	acquirer, err := r.acquirerRepo.GetAcquirer(id)
// 	if err != nil {
// 		return ErrNotFound // PATCH = update existing, don't create
// 	}

// 	// We assume that it's not possible to PATCH the ID, and that it's not
// 	// possible to PATCH any field to its zero value. That is, the zero value
// 	// means not specified. The way around this is to use e.g. Name *string in
// 	// the Acquirer definition.

// 	if a.Name != "" {
// 		acquirer.Name = a.Name
// 	}
// 	if len(a.Merchants) > 0 {
// 		acquirer.Merchants = a.Merchants
// 	}
// 	err = r.acquirerRepo.PatchAcquirer(id, a)
// 	return nil
// }

// func (r *acquirerService) DeleteAcquirer(id string) error {
// 	if err := validate.Validate(id); err != nil {
// 		return errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.PatchAcquirer")
// 	}
// 	if _, err := r.acquirerRepo.GetAcquirer(id); err != nil {
// 		return ErrNotFound
// 	}
// 	err := r.acquirerRepo.DeleteAcquirer(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *acquirerService) GetMerchant(acquirerID string, merchantID string) (Merchant, error) {
// 	if err := validate.Validate(acquirerID); err != nil {
// 		return Merchant{}, errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.GetMerchant")
// 	}
// 	if err := validate.Validate(merchantID); err != nil {
// 		return Merchant{}, errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.GetMerchant")
// 	}
// 	merchants, err := r.acquirerRepo.GetMerchants(acquirerID)
// 	if err != nil {
// 		return Merchant{}, ErrNotFound
// 	}
// 	for _, merchant := range merchants {
// 		if merchant.ID == merchantID {
// 			return merchant, nil
// 		}
// 	}
// 	return Merchant{}, ErrNotFound
// }

// func (r *acquirerService) PostMerchant(acquirerID string, m Merchant) error {
// 	if err := validate.Validate(acquirerID); err != nil {
// 		return errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.PostMerchant")
// 	}
// 	if err := validate.Validate(m); err != nil {
// 		return errs.Wrap(ErrMerchantInvalid, "service.Acquirer.PostMerchant")
// 	}

// 	merchants, err := r.acquirerRepo.GetMerchants(acquirerID)

// 	if err != nil {
// 		return ErrNotFound
// 	}
// 	for _, merchant := range merchants {
// 		if merchant.ID == m.ID {
// 			return ErrAlreadyExists
// 		}
// 	}
// 	err = r.acquirerRepo.PostMerchant(acquirerID, m)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *acquirerService) DeleteMerchant(acquirerID string, merchantID string) error {
// 	if err := validate.Validate(acquirerID); err != nil {
// 		return errs.Wrap(ErrAcquirerIdInvalid, "service.Acquirer.DeleteMerchant")
// 	}
// 	if err := validate.Validate(merchantID); err != nil {
// 		return errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.DeleteMerchant")
// 	}

// 	merchants, err := r.acquirerRepo.GetMerchants(acquirerID)
// 	if err != nil {
// 		return ErrNotFound
// 	}

// 	for _, merchant := range merchants {
// 		if merchant.ID == merchantID {
// 			continue // delete
// 		}

// 	}
// 	err = r.acquirerRepo.DeleteMerchant(acquirerID, merchantID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *acquirerService) GetTransaction(merchantID string, transactionID string) (Transaction, error) {
// 	if err := validate.Validate(merchantID); err != nil {
// 		return Transaction{}, errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.GetTransaction")
// 	}
// 	if err := validate.Validate(transactionID); err != nil {
// 		return Transaction{}, errs.Wrap(ErrTransactionIdInvalid, "service.Acquirer.GetTransaction")
// 	}

// 	transactions, err := r.acquirerRepo.GetTransactions(merchantID)
// 	if err != nil {
// 		return Transaction{}, ErrNotFound
// 	}
// 	for _, transaction := range transactions {
// 		if transaction.ID == transactionID {
// 			return Transaction{}, nil
// 		}
// 	}
// 	return Transaction{}, ErrNotFound
// }

// func (r *acquirerService) PostTransaction(merchantID string, t Transaction) error {
// 	if err := validate.Validate(merchantID); err != nil {
// 		return errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.PostTransaction")
// 	}
// 	if err := validate.Validate(t); err != nil {
// 		return errs.Wrap(ErrTransactionInvalid, "service.Acquirer.PostTransaction")
// 	}

// 	transactions, err := r.acquirerRepo.GetTransactions(merchantID)

// 	if err != nil {
// 		return ErrNotFound
// 	}

// 	for _, transaction := range transactions {
// 		if transaction.ID == t.ID {
// 			return ErrAlreadyExists
// 		}
// 	}
// 	err = r.acquirerRepo.PostTransaction(merchantID, t)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *acquirerService) DeleteTransaction(merchantID string, transactionID string) error {
// 	if err := validate.Validate(merchantID); err != nil {
// 		return errs.Wrap(ErrMerchantIdInvalid, "service.Acquirer.DeleteTransaction")
// 	}
// 	if err := validate.Validate(transactionID); err != nil {
// 		return errs.Wrap(ErrTransactionIdInvalid, "service.Acquirer.DeleteTransaction")
// 	}

// 	transacions, err := r.acquirerRepo.GetTransactions(merchantID)
// 	if err != nil {
// 		return ErrNotFound
// 	}

// 	for _, transaction := range transacions {
// 		if transaction.ID == transactionID {
// 			continue // delete
// 		}

// 	}

// 	err = r.acquirerRepo.DeleteTransaction(merchantID, transactionID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
