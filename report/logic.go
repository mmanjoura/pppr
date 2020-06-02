package report

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mmanjoura/pppr/configuration/decoder"
	"github.com/mmanjoura/pppr/logging"
	loggingpb "github.com/mmanjoura/pppr/logging/proto"
	errs "github.com/pkg/errors"
	"google.golang.org/grpc"

	. "github.com/ahmetb/go-linq"
	"gopkg.in/dealancer/validate.v2"
)

var (
	// ErrTransactionNotFound ...
	ErrTransactionNotFound = errs.New("Transaction Not Found")

	// ErrTransactionInvalid ...
	ErrTransactionInvalid = errors.New("Transaction Invaid")

	// ErrReportInvialid ...
	ErrReportInvalid = errors.New("Report Invalid")
	loggingClient    loggingpb.LogServiceClient
	requestCtx       context.Context
	requestOpts      grpc.DialOption
)

type reportService struct {
	reportRepo Repository
}

// NewReportService ...
func NewReportService(reportRepo Repository) Service {
	return &reportService{
		reportRepo,
	}
}

func (r *reportService) GetPayments(acquirerId, collection string) ([]Payment, map[string]string, error) {
	if err := validate.Validate(acquirerId); err != nil {
		return []Payment{}, nil, errs.Wrap(ErrReportInvalid, "service.Report.Transaction.Generate")
	}

	payments, mids, err := r.reportRepo.GetPayments(acquirerId, collection)
	if err != nil {

		return []Payment{}, nil, err
	}

	return payments, mids, nil
}

func (r *reportService) GenerateReports(reports []Report, acquirerId, collection string) error {

	payments, mids, err := r.reportRepo.GetPayments(acquirerId, collection)
	if err != nil {
		return err
	}

	reportConfig, err := decoder.Decode(decoder.ReportService)
	var hdr = []string{"Transaction ID", "Merchant ID", "Terminal ID", "CardNumber", "Amount", "Currency", "TxrDate", "AFAmount", "AFRate", "CardPresent"}
	var gPayments []Payment
	for _, v := range mids {

		From(payments).Where(
			func(p interface{}) bool {
				return p.(Payment).MerchantID == v
			}).Select(func(p interface{}) interface{} {
			return p.(Payment)
		}).ToSlice(&gPayments)

		pdfReport := GeneratePdf(gPayments, hdr, reportConfig, v)
		reports = append(reports, pdfReport)
		xlsReport := GenerateXls(gPayments, hdr, reportConfig, v)
		reports = append(reports, xlsReport)
		csvReport := GenerateCsv(gPayments, hdr, reportConfig, v)
		reports = append(reports, csvReport)

	}

	err = r.reportRepo.GenerateReports(reports, acquirerId, collection)
	if err != nil {
		return err
	}

	return nil
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
