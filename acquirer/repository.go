package acquirer

type Repository interface {

	// PostAcquirer(a Acquirer) error
	// GetAcquirer(id string) (Acquirer, error)
	// PutAcquirer(id string, a Acquirer) error
	// PatchAcquirer(id string, a Acquirer) error
	// DeleteAcquirer(id string) error

	GetAcquirerPayments(acquirerID string) ([]Payment, error)
	GetMerchantPayments(MID string) ([]Payment, error)
	GetReports(acquirerID string) ([]Report, error)
	GetExchangeRates(date string) (ExchangeRate, error)
	GetLogMessages(date string) ([]Message, error)

	GetProcessStates(approved bool) ([]ProcessState, error)
	PutProcessState(processID string, ps ProcessState) error
	// GetMerchant(acquirerID string, merchantID string) (Merchant, error)
	// PostMerchant(acquirerID string, a Merchant) error
	// DeleteMerchant(acquirerID string, merchantID string) error

	GetTransactions(merchantID string) ([]Transaction, error)
	// GetTransaction(merchantID string, transactionID string) (Transaction, error)
	// PostTransaction(merchantID string, a Transaction) error
	// DeleteTransaction(merchantID string, transaction string) error
}
