package report

// Repository
type Repository interface {
	GetPayments(acquierId, collection string) ([]Payment, map[string]string, error)
	// GenerateReports(reports []Report, collection string) error
	GenerateReports(reports []Report, acquirerId, collection string) error
}
