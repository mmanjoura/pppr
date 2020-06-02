package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"
	"github.com/mmanjoura/pppr/payment"
	"github.com/mmanjoura/pppr/report"
	reportpb "github.com/mmanjoura/pppr/report/proto"
	"github.com/mmanjoura/pppr/report/repository/mongo"

	"google.golang.org/grpc"
)

// ReportServiceServer ...
type ReportServiceServer struct {
}

var (
	hotName string
)

// GenerateReport ...
func (s *ReportServiceServer) GenerateReport(ctx context.Context, req *reportpb.GenerateReportReq) (*reportpb.GenerateReportRes, error) {
	hotName, err := os.Hostname()
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)
	trxConfig, err := decoder.Decode(decoder.ReportService)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Report Service",
			CallingMethod: "GenerateReport",
			Host:          hotName,
			Body:          "Failed to Run Report " + fmt.Sprint(err.Error()),
			Latency:       "NA",
		}
		err := report.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}

	repo := chooseRepo(trxConfig)
	reportService := report.NewReportService(repo)
	reportMeta := report.Report{}
	reportReq := req.GetReport()

	reportMeta.AcquirerID = reportReq.GetAcquirerID()
	// reportMeta.StartDate = reportReq.GetStartDate()
	// reportMeta.EndDate = reportReq.GetEndDate()
	reportMeta.ReportType = reportReq.GetReportType()
	fmt.Printf("Received Request to generate reports for AcquirerID: %s at: %s \n", reportMeta.AcquirerID, time.Now().Format(time.RFC3339))
	dbLogger := logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Report Service",
		CallingMethod: "GenerateReport",
		Host:          hotName,
		Body:          "Received Request to Generate Report for AcquirerID:" + reportMeta.AcquirerID,
		Latency:       "NA",
	}
	err = report.LogMessage(dbLogger, loggingSvcConfig)
	if err != nil {
		log.Fatalf("Could not log Message: %v", err)
	}

	fmt.Printf("Generating  reports for AcquirerID: %s at: %s \n", reportMeta.AcquirerID, time.Now().Format(time.RFC3339))
	dbLogger = logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Report Service",
		CallingMethod: "GenerateReport",
		Host:          hotName,
		Body:          "Received Request to Generate Report for AcquirerID:" + reportMeta.AcquirerID,
		Latency:       "NA",
	}

	start := time.Now()
	err = reportService.GenerateReports([]report.Report{}, reportReq.GetAcquirerID(), trxConfig.CollectionName)
	end := time.Since(start)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Report Service",
			CallingMethod: "GenerateReport",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       end.String(),
		}
		err := report.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}
	fmt.Printf("Report Generated for AcquirerID: %s at: %s \n", reportMeta.AcquirerID, time.Now().Format(time.RFC3339))

	dbLogger = logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Report Service",
		CallingMethod: "GenerateReport",
		Host:          hotName,
		Body:          "Reports Generated Sucessfully",
		Latency:       end.String(),
	}
	err = report.LogMessage(dbLogger, loggingSvcConfig)
	if err != nil {
		log.Fatalf("Could not log Message: %v", err)
	}

	return &reportpb.GenerateReportRes{Report: reportReq}, nil

}

func main() {

	hotName, err := os.Hostname()
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)
	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Report Service",
			CallingMethod: "main.main",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}

		err := payment.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}
	reportSvcConfig, err := decoder.Decode(decoder.ReportService)
	gRPCreportServer := fmt.Sprintf("%s:%d", reportSvcConfig.GRPCServer, reportSvcConfig.GRPCPort)
	listener, err := net.Listen("tcp", gRPCreportServer)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Report Service",
			CallingMethod: "main.main",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}

		err := payment.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &ReportServiceServer{}
	reportpb.RegisterReportServiceServer(s, srv)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %s %v", gRPCreportServer, err)
		}
	}()

	fmt.Printf("Running Report Service on %s Port: %v \n", hotName, reportSvcConfig.GRPCPort)

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Done.")
}

func chooseRepo(rc decoder.Config) report.Repository {

	switch rc.DbHost {
	case "mongo":
		mongoURL := rc.Mongo.Mongourl
		mongodb := rc.DbName
		mongoTimeout := rc.Mongo.MongoTimeout
		repo, err := mongo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
