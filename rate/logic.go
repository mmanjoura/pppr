package rate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"
	loggingpb "github.com/mmanjoura/pppr/logging/proto"
	"google.golang.org/grpc"

	finance "github.com/pieterclaerhout/go-finance"
	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrRatesInvalid ...
	ErrRatesInvalid = errors.New("Rate Invalid")

	// loggingClient
	loggingClient loggingpb.LogServiceClient

	requestCtx  context.Context
	requestOpts grpc.DialOption
)

type service struct {
	exchangeRatesRepo Repository
}

// NewService ...
func NewService(exchangeRatesRepo Repository) Service {
	return &service{
		exchangeRatesRepo,
	}
}

// Save ...
func (s *service) Save(rates ExchangeRateTimedCube, collection string) error {

	if err := validate.Validate(rates); err != nil {
		return errs.Wrap(ErrRatesInvalid, "service.drg.Save")
	}

	rateConfig, err := decoder.Decode(decoder.DrgService)
	exchangeRateTimedCube, err := getRates()

	err = s.exchangeRatesRepo.Save(exchangeRateTimedCube, rateConfig.CollectionName)

	if err != nil {
		return err
	}

	return nil
}

func getRates() (ExchangeRateTimedCube, error) {

	ex := ExchangeRateCurrencyCube{}
	ext := ExchangeRateTimedCube{}
	ext.CreatedDate = time.Now().Format("2006-01-02")
	ext.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[10:]

	rates, err := finance.ExchangeRates()
	if err != nil {
		return ExchangeRateTimedCube{}, err
	}
	for currency, rate := range rates {

		ex.Currency = currency
		ex.Rate = rate
		ext.Rates = append(ext.Rates, ex)

	}

	return ext, nil
}

// LogMessage ...
func LogMessage(srvlog logging.LogMessage, conf decoder.Config) error {

	Log := &loggingpb.Log{}

	// Prepare connection to Log gRP Server
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	gRPCServer := fmt.Sprintf("%s:%d", conf.GRPCServer, conf.GRPCPort)
	conn, err := grpc.Dial(gRPCServer, requestOpts)
	if err != nil {
		conn.Close()
		return err
	}

	loggingClient = loggingpb.NewLogServiceClient(conn)

	Log.Level = srvlog.Level
	Log.CallingMethod = srvlog.CallingMethod
	Log.ServiceName = srvlog.ServiceName
	Log.Host = srvlog.Host
	Log.Body = srvlog.Body
	Log.Latency = srvlog.Latency

	_, err = loggingClient.SaveLog(
		context.TODO(),
		&loggingpb.SaveLogReq{
			Log: Log,
		},
	)

	if err != nil {
		conn.Close()
		return err
	}
	conn.Close()
	return nil
}
