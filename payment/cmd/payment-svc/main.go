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
	paymentpb "github.com/mmanjoura/pppr/payment/proto"
	"github.com/mmanjoura/pppr/payment/repository/mongo"

	"google.golang.org/grpc"
)

// PaymentServiceServer ...
type PaymentServiceServer struct {
}

var (
	hotName string
)

func main() {

	hotName, err := os.Hostname()
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Payment Service",
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

	paymentSvcConfig, err := decoder.Decode(decoder.PaymentService)
	gRPCpaymentServer := fmt.Sprintf("%s:%d", paymentSvcConfig.GRPCServer, paymentSvcConfig.GRPCPort)
	listener, err := net.Listen("tcp", gRPCpaymentServer)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Payment Service",
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
	srv := &PaymentServiceServer{}

	paymentpb.RegisterPaymentServiceServer(s, srv)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %s %v", gRPCpaymentServer, err)
		}
	}()

	fmt.Printf("Running Payment Service on %s Port: %v \n", hotName, paymentSvcConfig.GRPCPort)

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

	fmt.Println("Received Request to Run Payment at:", time.Now().Format(time.RFC3339))
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)
	config, err := decoder.Decode(decoder.PaymentService)
	hotName, err := os.Hostname()

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Payment Service",
			CallingMethod: "RunPayment",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}
		err := payment.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
		return &paymentpb.RunPaymentRes{}, err
	}

	repo := chooseRepo(config)
	paymentService := payment.NewPaymentService(repo)
	pp := payment.RunParams{}
	pbPayment := req.GetPayment()

	pp.AcquirerID = pbPayment.GetAcquirerId()
	pp.StartDate = pbPayment.GetStartDate()
	pp.EndDate = pbPayment.GetStartDate()

	pAcquirer, err := GetPAcquirer(config, paymentService, pp.AcquirerID)

	if err != nil {

		return &paymentpb.RunPaymentRes{}, err
	}

	dbLogger := logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Payment Service",
		CallingMethod: "RunPayment",
		Host:          hotName,
		Body:          "Payment Run Statred for : " + pp.AcquirerID,
		Latency:       "NA",
	}
	err = payment.LogMessage(dbLogger, loggingSvcConfig)
	if err != nil {
		log.Fatalf("Could not log Message: %v", err)
	}
	fmt.Println("Payment Run Started at:", time.Now().Format(time.RFC3339))
	start := time.Now()
	err = paymentService.Run(&pp, config.CollectionName, &pAcquirer)
	end := time.Since(start)
	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Payment Service",
			CallingMethod: "RunPayment",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       end.String(),
		}
		err := payment.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
		return &paymentpb.RunPaymentRes{}, err
	}
	dbLogger = logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Payment Service",
		CallingMethod: "RunPayment",
		Host:          hotName,
		Body:          "Payments Run finished Sucessfully for: " + pp.AcquirerID,
		Latency:       end.String(),
	}
	err = payment.LogMessage(dbLogger, loggingSvcConfig)
	if err != nil {
		log.Fatalf("Could not log Message: %v", err)
	}
	fmt.Println("Payment Run finished at:", time.Now().Format(time.RFC3339))
	// Now you can run Report
	//Now run reports, we need to return an error here
	reportParams := payment.Report{}
	//reportParams.StartDate = pp.StartDate
	//reportParams.EndDate = pp.EndDate
	reportParams.AcquirerID = pp.AcquirerID
	reportParams.ReportType = "Not Known"

	reportSvcConfig, err := decoder.Decode(decoder.ReportService)
	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Payment Service",
			CallingMethod: "RunPayment",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}

		err := payment.LogMessage(dbLogger, loggingSvcConfig)

		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}

	dbLogger = logging.LogMessage{
		Level:         "INFO",
		ServiceName:   "Payment Service",
		CallingMethod: "RunPayment",
		Host:          hotName,
		Body:          "Calling Generate Reports in Report Service: ",
		Latency:       "NA",
	}
	err = payment.LogMessage(dbLogger, loggingSvcConfig)
	if err != nil {
		log.Fatalf("Could not log Message: %v", err)
	}
	fmt.Println("Calling Report Service to Generate Reports at:", time.Now().Format(time.RFC3339))
	start = time.Now()
	err = payment.GenerateReport(reportParams, reportSvcConfig)
	end = time.Since(start)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Report Service",
			CallingMethod: "RunPayment",
			Host:          hotName,
			Body:          "Failed to Run Report through gRPC server: " + fmt.Sprint(err.Error()),
			Latency:       end.String(),
		}
		err := payment.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}

	fmt.Println("Generate Report : OK, for acquirerID: ", pp.AcquirerID)

	return &paymentpb.RunPaymentRes{Payment: pbPayment}, nil

}

// GetPAcquirer ...
func GetPAcquirer(config decoder.Config, paymentService payment.Service, id string) (payment.PAcquirer, error) {
	acquirer := payment.Acquirer{}
	merchant := payment.Merchant{}
	priceOption := payment.PriceOption{}

	pAcquirer := payment.PAcquirer{}
	pMerchant := payment.PMerchant{}
	pOption := payment.POption{}

	pMerchants := []payment.PMerchant{}
	pOptions := []payment.POption{}

	model, err := paymentService.Get(acquirer, config.MmsBaseURL+config.MmsEndPoint)
	if err != nil {
		return payment.PAcquirer{}, err
	}

	// "a" here is the acquirer Model
	if acq, ok := model.(payment.Acquirer); ok {
		for _, a := range acq.Data {
			if a.AquirerID != id {
				continue
			}
			pAcquirer.UUID = a.UUID
			pAcquirer.Name = a.Name
			pAcquirer.AquirerID = a.AquirerID

			// Iterate through the merchant URL
			for j, mrul := range a.Merchant {

				// Get the merchant
				mer, err := paymentService.Get(merchant, config.MmsBaseURL+mrul)
				if err != nil {
					return payment.PAcquirer{}, err
				}

				if m, ok := mer.(payment.Merchant); ok {

					pMerchant.UUID = m.Data[j].UUID
					pMerchant.Name = m.Data[j].Name
					pMerchant.Address = m.Data[j].Address
					pMerchant.Mid = m.Data[j].Mid
					pMerchant.Tid = m.Data[j].Tid

					// Get Price Option Json, notice we have
					// only one Price option per merchant hence "0"

					opt, err := paymentService.Get(priceOption, config.MmsBaseURL+m.Data[j].Option)
					if err != nil {
						return payment.PAcquirer{}, err
					}
					if o, ok := opt.(payment.PriceOption); ok {
						for _, po := range o.Data {

							pOption.UUID = po.UUID
							pOption.Name = po.Name
							pOption.Type = po.Type
							pOption.Scheme = po.Scheme
							pOption.DomesticMSCPPT = po.DomesticMSCPPT
							pOption.EEAMSCPPT = po.EEAMSCPPT
							pOption.EEAMSCRate = po.EEAMSCRate
							pOption.MSCRate = po.MSCRate

							pOptions = append(pOptions, pOption)

						}
					}
				}
			}
			pMerchant.POptions = pOptions
			pMerchants = append(pMerchants, pMerchant)
			pAcquirer.PMerchants = pMerchants
		}

	}

	return pAcquirer, nil
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
