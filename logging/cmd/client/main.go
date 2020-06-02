package main

import (
	"context"
	"log"
	"time"

	logpb "github.com/mmanjoura/pppr/logging/proto"
	"google.golang.org/grpc"
)

var client logpb.LogServiceClient
var requestCtx context.Context
var requestOpts grpc.DialOption

func main() {

	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50050", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}

	client = logpb.NewLogServiceClient(conn)

	logging := &logpb.Log{

		CreatedDate:   "2020-04-04 00:00:00",
		CreatedTime:   "2020-04-04 00:00:00",
		Level:         "Error",
		ServiceName:   "Logging",
		CallingMethod: "Main.main",
		Host:          "Localhost",
		Body:          "This is the description of the message",
		Latency:       "34ms",
	}
	_, err = client.SaveLog(
		context.TODO(),
		&logpb.SaveLogReq{
			Log: logging,
		},
	)
	if err != nil {
		log.Fatalf("Could not create Logging: %v", err)
	}
}
