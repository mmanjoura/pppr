package transaction

type Service interface {
	Save(meta *Meta, collection string) (map[string]string, error)
	Get(date, time string) ([]string, error)
}
