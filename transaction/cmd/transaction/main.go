package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mmanjoura/pppr/configuration/decoder"
	"github.com/mmanjoura/pppr/logging"
	"github.com/mmanjoura/pppr/transaction"
	"github.com/mmanjoura/pppr/transaction/repository/mongo"
)

// repo <- service -> serializer  -> transport

var (
	hostname string
)

func main() {
	hostname, err := os.Hostname()

	// Get Transaction Service configuration element
	// than Start listening for TXF response files
	trxSvcconfig, err := decoder.Decode(decoder.TransactionService)
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)
	paymentSvcConfig, err := decoder.Decode(decoder.PaymentService)

	// Report the errors above
	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "Transaction Service",
			CallingMethod: "main.main",
			Host:          hostname,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}
		err := transaction.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Could not log Message: %v", err)
		}
	}

	repo := chooseRepo(trxSvcconfig)
	transactionService := transaction.NewService(repo)

	// Loop for scan inerval looking for files
	fmt.Printf("running Transaction Service on %s Port: %v \n", hostname, trxSvcconfig.HTTPPort)
	go func() {
		for {

			meta := transaction.Meta{}
			start := time.Now()
			dateTime, err := transactionService.Save(&meta, trxSvcconfig.CollectionName)
			end := time.Since(start)
			if err != nil {
				dbLogger := logging.LogMessage{
					Level:         "ERROR",
					ServiceName:   "Transaction Service",
					CallingMethod: "main.main",
					Host:          hostname,
					Body:          fmt.Sprint(err),
					Latency:       end.String(),
				}
				transaction.LogMessage(dbLogger, loggingSvcConfig)
				if err != nil {
					log.Fatalf("Error could not log to MongoDb %v", err)
				}
			}

			// For each unique dateTime run payment
			for t, d := range dateTime {
				fmt.Printf("date map: %s", dateTime)
				runParams := transaction.RunParams{}
				if len(d) > 0 {

					dbLogger := logging.LogMessage{
						Level:         "INFO",
						ServiceName:   "Transaction Service",
						CallingMethod: "main.main",
						Host:          hostname,
						Body:          "Imported TXF file successful",
						Latency:       end.String(),
					}
					err = transaction.LogMessage(dbLogger, loggingSvcConfig)
					if err != nil {
						log.Fatalf("Error could not log to MongoDb %v", err)
					}

					fmt.Println("Imported TXF files successful at: ", time.Now().Format(time.RFC3339))

					fmt.Println(d)
					fmt.Println(t)

					runParams.CreatedDate = d
					runParams.CreatedTime = t

					acquirerIDs, err := transactionService.Get(d, t)

					// Run Payments for each acquirer
					for _, acqID := range acquirerIDs {
						if err != nil {
							log.Fatalf("could not get acquirers %v", err)
						}
						runParams.AcquirerID = acqID

						fmt.Printf("acquirers IDS: %s", acqID)

						fmt.Println("Calling Payment RPC server at: ", time.Now().Format(time.RFC3339))

						dbLogger = logging.LogMessage{
							Level:         "INFO",
							ServiceName:   "Transaction Service",
							CallingMethod: "main.main",
							Host:          hostname,
							Body:          "Calling Payment Service To run Payment",
							Latency:       "NA",
						}
						err = transaction.LogMessage(dbLogger, loggingSvcConfig)
						if err != nil {
							log.Fatal("Error could not log to MongoDb", err)
						}

						start := time.Now()
						err = transaction.RunPayment(runParams, paymentSvcConfig)
						end := time.Since(start)
						if err != nil {
							dbLogger := logging.LogMessage{
								Level:         "ERROR",
								ServiceName:   "Payment Service",
								CallingMethod: "main.main",
								Host:          hostname,
								Body:          "Failed to Run Payment through gRPC server: " + fmt.Sprint(err.Error()),
								Latency:       end.String(),
							}
							err := transaction.LogMessage(dbLogger, loggingSvcConfig)
							if err != nil {
								log.Fatalf("Could not log Message: %v", err)
							}
						}

					}

					fmt.Println("Received Successfull Response from Payment Service at:", time.Now().Format(time.RFC3339))

				}

			}

			time.Sleep(time.Second * time.Duration(trxSvcconfig.ScanInterval))
		}
	}()

	// Watch for any interrupt like Ctr-c
	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}

func chooseRepo(conf decoder.Config) transaction.Repository {

	switch conf.DbHost {

	case "mongo":
		mongoURL := conf.Mongo.Mongourl
		mongodb := conf.DbName
		mongoTimeout := conf.Mongo.MongoTimeout
		repo, err := mongo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
