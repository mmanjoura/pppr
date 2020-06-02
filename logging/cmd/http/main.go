package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"
	"github.com/mmanjoura/pppr/logging/api"
	"github.com/mmanjoura/pppr/logging/repository/mongo"
)

// repo <- service -> serializer  -> http

func main() {

	trxConfig, err := decoder.Decode(decoder.LoggingService)
	if err != nil {
		log.Fatal(err)
	}
	repo := chooseRepo(trxConfig)
	service := logging.NewLoggingService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/{collection}", handler.Post)

	errs := make(chan error, 2)
	go func() {

		fmt.Println("Listening on port : " + strconv.Itoa(trxConfig.HTTPPort) + " started at: " + time.Now().Format(time.RFC3339))
		errs <- http.ListenAndServe(":"+strconv.Itoa(trxConfig.HTTPPort), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func chooseRepo(rc decoder.Config) logging.LoggingRepository {

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
