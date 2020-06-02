package main

import (
	"context"
	"fmt"
	"log"
	"time"

	reportpb "github.com/mmanjoura/pppr/report/proto"
	"google.golang.org/grpc"
)

// Use this Client to test the Report gRPC server
var client reportpb.ReportServiceClient
var requestCtx context.Context
var requestOpts grpc.DialOption

func main() {

	fmt.Println("Starting Report Service Client")

	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	requestOpts = grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50080", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50080: %v", err)
	}

	client = reportpb.NewReportServiceClient(conn)

	report := &reportpb.Report{

		AcquirerID: "d41035e2-b1bc-4bfa-a11f-de842c5b69ad",
		//StartDate:  "2020-04-06 10:15:49",
		//EndDate:    "2020-04-06 10:15:49",
		ReportType: "Pdf",
	}
	_, err = client.GenerateReport(
		context.TODO(),
		&reportpb.GenerateReportReq{
			Report: report,
		},
	)
	if err != nil {
		log.Fatalf("Could not create report: %v", err)
	}
}
