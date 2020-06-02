package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/mmanjoura/pppr/configuration/decoder"
	"github.com/mmanjoura/pppr/payment"
	paymentpb "github.com/mmanjoura/pppr/payment/proto"
	"github.com/mmanjoura/pppr/payment/repository/mongo"

	"google.golang.org/grpc"
)

// PaymentServiceServer ...
type PaymentServiceServer struct {
}

func main() {
	paymentSvcConfig, err := decoder.Decode(decoder.PaymentService)
	gRPCpaymentServer := fmt.Sprintf("%s:%d", paymentSvcConfig.GRPCServer, paymentSvcConfig.GRPCPort)
	listener, err := net.Listen("tcp", gRPCpaymentServer)

	if err != nil {
		log.Fatalf("Unable to listen on %s %v", gRPCpaymentServer, err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := &PaymentServiceServer{}
	paymentpb.RegisterPaymentServiceServer(s, srv)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %s %v", gRPCpaymentServer, err)
		}
	}()
	fmt.Println("Server succesfully started on %s, at: %s ", gRPCpaymentServer, time.Now().Format(time.RFC3339))

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Done.")
}

//RunPayment ...
func (s *PaymentServiceServer) RunPayment(ctx context.Context, req *paymentpb.RunPaymentReq) (*paymentpb.RunPaymentRes, error) {

	config, err := decoder.Decode(decoder.PaymentService)
	if err != nil {
		log.Fatal(err)
	}

	repo := chooseRepo(config)
	paymentService := payment.NewPaymentService(repo)
	pp := payment.RunParams{}
	paymentPb := req.GetPayment()

	pp.AcquirerID = paymentPb.GetAcquirerId()
	pp.StartDate = paymentPb.GetStartDate()
	pp.EndDate = paymentPb.GetStartDate()

	paymentAcquirer := payment.PAcquirer{}

	err = paymentService.Run(&pp, config.CollectionName, &paymentAcquirer)
	if err != nil {
		log.Fatal(err)
	}
	return &paymentpb.RunPaymentRes{Payment: paymentPb}, nil

}
func chooseRepo(rc decoder.Config) payment.Repository {

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
