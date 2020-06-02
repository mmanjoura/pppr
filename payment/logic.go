package payment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"
	loggingpb "github.com/mmanjoura/pppr/logging/proto"
	reportpb "github.com/mmanjoura/pppr/report/proto"
	errs "github.com/pkg/errors"
	"google.golang.org/grpc"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrTransactionNotFound ...
	ErrTransactionNotFound = errors.New("Transaction Not Found")

	ErrMerchantNotFound = errors.New("Merchant Not Found")

	// ErrTransactionInvalid ..
	ErrTransactionInvalid = errors.New("Transaction Invalid")

	reportClient  reportpb.ReportServiceClient
	loggingClient loggingpb.LogServiceClient
	requestCtx    context.Context
	requestOpts   grpc.DialOption
)

type paymentService struct {
	paymentRepo Repository
}

// NewPaymentService ...
func NewPaymentService(paymentRepo Repository) Service {
	return &paymentService{
		paymentRepo,
	}
}

// Run Payment for the given run params
func (r *paymentService) Run(runParams *RunParams, collection string, model *PAcquirer) error {
	if err := validate.Validate(runParams); err != nil {
		return errs.Wrap(ErrTransactionInvalid, "service.Payment.Transaction.Run")
	}

	return r.paymentRepo.Run(runParams, collection, model)
}

func (r *paymentService) Get(model interface{}, url string) (interface{}, error) {
	if err := validate.Validate(model); err != nil {
		return nil, errs.Wrap(ErrTransactionInvalid, "service.Payment.Transaction.Run")
	}

	client := new(http.Client)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//request.Header.Add("Accept-Encoding", "gzip")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Interfaces are magic in Go, look at this call
	// for different routes, without knowing which one
	switch model.(type) {

	case Acquirer:
		resp := Acquirer{}
		err = json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case Merchant:
		resp := Merchant{}
		err = json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case PriceOption:
		resp := PriceOption{}
		err = json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return nil, err
		}
		return resp, nil

	default:
		return "Not Known model Type", nil
	}

}

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

func GenerateReport(report Report, conf decoder.Config) error {

	rep := &reportpb.Report{}

	// Prepare connection to Log gRP Server
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	gRPCServer := fmt.Sprintf("%s:%d", conf.GRPCServer, conf.GRPCPort)
	conn, err := grpc.Dial(gRPCServer, requestOpts)
	if err != nil {
		conn.Close()
		return err
	}

	reportClient = reportpb.NewReportServiceClient(conn)

	rep.AcquirerID = report.AcquirerID
	// rep.StartDate = report.StartDate
	// rep.EndDate = report.EndDate
	rep.ReportType = report.ReportType

	_, err = reportClient.GenerateReport(
		context.TODO(),
		&reportpb.GenerateReportReq{
			Report: rep,
		},
	)

	if err != nil {
		conn.Close()
		return err
	}
	conn.Close()
	return nil
}
