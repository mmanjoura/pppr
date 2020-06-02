package logging

type LoggingRepository interface {
	Save(logging *LogMessage, collection string) error
}
