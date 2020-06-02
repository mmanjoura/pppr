package payment

//Service ...
type Service interface {
	Run(runParams *RunParams, collection string, model *PAcquirer) error
	Get(model interface{}, url string) (interface{}, error)
}
