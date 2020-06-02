package logging

import (
	"context"
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"google.golang.org/grpc"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrLoggingNotFound ...
	ErrLoggingNotFound = errors.New("Logging Not Found")

	// ErrLoggingInvalid ...
	ErrLoggingInvalid = errors.New("Logging Invalid")

	//loggingClient loggingpb.LogServiceClient
	requestCtx  context.Context
	requestOpts grpc.DialOption
)

type loggingService struct {
	loggingRepo LoggingRepository
}

// NewLoggingService ...
func NewLoggingService(loggingRepo LoggingRepository) LoggingService {
	return &loggingService{
		loggingRepo,
	}
}

// Save ...
func (r *loggingService) Save(logging *LogMessage, collection string) error {
	if err := validate.Validate(logging); err != nil {
		return errs.Wrap(ErrLoggingInvalid, "service.Logging.Save")
	}
	logging.CreatedDate = time.Now().Format("2006-01-02")
	logging.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[10:]

	return r.loggingRepo.Save(logging, collection)
}
