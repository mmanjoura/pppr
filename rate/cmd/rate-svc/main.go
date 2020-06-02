package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mmanjoura/pppr/rate"

	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"
	"github.com/mmanjoura/pppr/rate/repository/mongo"
)

// repo <- service -> serializer  -> transport

var (
	hotName string
)

func main() {
	hotName, err := os.Hostname()
	if err != nil {
		log.Fatal("Failed to start DRG Service, Cannot get host name")
	}

	rateSvcconfig, err := decoder.Decode(decoder.DrgService)
	loggingSvcConfig, err := decoder.Decode(decoder.LoggingService)

	if err != nil {
		dbLogger := logging.LogMessage{
			Level:         "ERROR",
			ServiceName:   "drg Service",
			CallingMethod: "main.main",
			Host:          hotName,
			Body:          fmt.Sprint(err),
			Latency:       "NA",
		}
		err := rate.LogMessage(dbLogger, loggingSvcConfig)
		if err != nil {
			log.Fatalf("Error could not log to MongoDb %v", err)
		}
	}

	repo := chooseRepo(rateSvcconfig)
	drgService := rate.NewService(repo)

	go func() {
		for {
			rates := rate.ExchangeRateTimedCube{}
			start := time.Now()
			err := drgService.Save(rates, rateSvcconfig.CollectionName)
			end := time.Since(start)
			if err != nil {
				dbLogger := logging.LogMessage{
					Level:         "ERROR",
					ServiceName:   "DRG Service",
					CallingMethod: "main.main",
					Host:          hotName,
					Body:          fmt.Sprint(err),
					Latency:       end.String(),
				}

				rate.LogMessage(dbLogger, loggingSvcConfig)
				if err != nil {
					log.Fatalf("Could not log Message: %v", err)
				}
			}

			dbLogger := logging.LogMessage{
				Level:         "INFO",
				ServiceName:   "DRG Service",
				CallingMethod: "main.main",
				Host:          hotName,
				Body:          "Imported Rates Successfully",
				Latency:       end.String(),
			}

			// Log informtion for TXT import
			err = rate.LogMessage(dbLogger, loggingSvcConfig)
			if err != nil {
				log.Fatalf("Could not log Message: %v", err)
			}
			fmt.Println("Import Rates successfully from ECB at: ", time.Now().Format(time.RFC3339))
			time.Sleep(time.Second * time.Duration(rateSvcconfig.ScanInterval))
		}
	}()

	fmt.Printf("Running DRG Service on %s Port: %v \n", hotName, rateSvcconfig.HTTPPort)
	// Watch for any interrupt like Ctr-c
	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}

func chooseRepo(conf decoder.Config) rate.Repository {

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
