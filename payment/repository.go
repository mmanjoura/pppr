package payment

//Repository ...
type Repository interface {
	Run(runparams *RunParams, collection string, model *PAcquirer) error
}
