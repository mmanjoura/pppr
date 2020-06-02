package report

// Service ...
type Service interface {
	GetPayments(acquirerId, collection string) ([]Payment, map[string]string, error)
	//GenerateReports(reports []Report, collection string) error
	GenerateReports(reports []Report, acquirerId, collection string) error
}
