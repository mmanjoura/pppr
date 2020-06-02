package logging

type LoggingService interface {
	Save(logging *LogMessage, collection string) error
}
