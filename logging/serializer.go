package logging

type LoggingSerializer interface {
	Decode(input []byte) (*LogMessage, error)
	Encode(input *LogMessage) ([]byte, error)
}
