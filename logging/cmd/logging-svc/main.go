package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/mmanjoura/pppr/configuration/decoder"
	"github.com/mmanjoura/pppr/logging"
	loggingpb "github.com/mmanjoura/pppr/logging/proto"
	"github.com/mmanjoura/pppr/logging/repository/mongo"
	"google.golang.org/grpc"
)

var (
	hotName string
)

// LogServiceServer ...
type LogServiceServer struct {
}

// SaveLog ...
func (s *LogServiceServer) SaveLog(ctx context.Context, req *loggingpb.SaveLogReq) (*loggingpb.SaveLogRes, error) {

	// Get the configuration params for Logging Service
	lgnConfig, err := decoder.Decode(decoder.LoggingService)
	if err != nil {
		log.Fatal(err)
	}

	// Get the type repository
	repo := getRepository(lgnConfig)
	loggingService := logging.NewLoggingService(repo)
	dbLogMessage := logging.LogMessage{}
	logging := req.GetLog()

	// Build the Logging message sent from clients
	dbLogMessage.Level = logging.GetLevel()
	dbLogMessage.ServiceName = logging.GetServiceName()
	dbLogMessage.CallingMethod = logging.GetCallingMethod()
	dbLogMessage.Host = logging.GetHost()
	dbLogMessage.Body = logging.GetBody()
	dbLogMessage.Latency = logging.GetLatency()

	// Save log message into the Db
	err = loggingService.Save(&dbLogMessage, lgnConfig.CollectionName)
	if err != nil {
		log.Fatal(err)
	}
	return &loggingpb.SaveLogRes{Log: logging}, nil
}

func main() {
	hotName, err := os.Hostname()
	if err != nil {
		log.Fatal("Failed to  start Logging Service, Cannot get host name")
	}
	// Get logging Config params
	loggingConfig, err := decoder.Decode(decoder.LoggingService)

	// Get Logging GRPS Server name and Port number
	gRPCLoggingServer := fmt.Sprintf("%s:%d", loggingConfig.GRPCServer, loggingConfig.GRPCPort)

	// Start Logging GRPC Server
	listener, err := net.Listen("tcp", gRPCLoggingServer)
	if err != nil {
		log.Fatalf("Unable to listen on %s %v", gRPCLoggingServer, err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &LogServiceServer{}

	loggingpb.RegisterLogServiceServer(s, srv)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", gRPCLoggingServer, err)
		}
	}()

	fmt.Printf("Running Logging Service on %s Port: %v \n", hotName, loggingConfig.GRPCPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Done.")

}

func getRepository(rc decoder.Config) logging.LoggingRepository {

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
