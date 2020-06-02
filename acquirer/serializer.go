package acquirer

// Service is a simple CRUD interface for an acquirers.
type Serializer interface {
	// EndcodePostAcquirer(input Acquirer) ([]byte, error)
	// EndcodeGetAcquirer(input Acquirer) ([]byte, error)
	// EndcodePutAcquirer(input Acquirer) ([]byte, error)
	// EndcodePatchAcquirer(input Acquirer) ([]byte, error)
	// EndcodeDeleteAcquirer(input Acquirer) ([]byte, error)

	// DecodePostAcquirer(input []byte) (Acquirer, error)
	// DecodeGetAcquirer(input []byte) (Acquirer, error)
	// DecodePutAcquirer(input []byte) (Acquirer, error)
	// DecodePatchAcquirer(input []byte) (Acquirer, error)
	// DecodeDeleteAcquirer(input []byte) (Acquirer, error)

	EncodeGetAcquirerPayments(input []Payment) ([]byte, error)
	EncodeGetMerchantPayments(input []Payment) ([]byte, error)
	EncodeGetReports(input []Report) ([]byte, error)
	EncodeGetExchangeRates(input ExchangeRate) ([]byte, error)
	EncodeGetLogMessages(input []Message) ([]byte, error)

	EncodeGetProcessStates(input []ProcessState) ([]byte, error)
	EncodePutProcessState(input ProcessState) error
	// EndcodeGetMerchant(input Merchant) ([]byte, error)
	// EndcodePostMerchant(input Merchant) ([]byte, error)
	// EndcodeDeleteMerchant(input Merchant) ([]byte, error)

	DecodeGetAcquirerPayments(input []byte) ([]Payment, error)
	DecodeGetMerchantPayments(input []byte) ([]Payment, error)
	DecodeGetReports(input []byte) ([]Report, error)
	DecodeGetExchangeRates(input []byte) (ExchangeRate, error)
	DecodeGetLogMessages(input []byte) ([]Message, error)
	// DecodeGetMerchant(input []byte) (Merchant, error)
	// DecodePostMerchant(input []byte) (Merchant, error)
	// DecodeDeleteMerchant(input []byte) (Merchant, error)

	EncodeGetTransactions(input []Transaction) ([]byte, error)
	// EndcodeGetTransaction(input Transaction) ([]byte, error)
	// EndcodePostTransaction(input Transaction) ([]byte, error)
	// EndcodeDeleteTransaction(input Transaction) ([]byte, error)

	DecodeGetTransactions(input []byte) ([]Transaction, error)
	// DecodeGetTransaction(input []byte) (Transaction, error)
	// DecodePostTransaction(input []byte) (Transaction, error)
	// DecodeDeleteTransaction(input []byte) (Transaction, error)
}
