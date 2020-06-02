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
	"github.com/go-chi/cors"
	"github.com/mmanjoura/pppr/acquirer"
	"github.com/mmanjoura/pppr/acquirer/api"
	"github.com/mmanjoura/pppr/acquirer/repository/mongo"
	"github.com/mmanjoura/pppr/configuration/decoder"
)

func main() {

	trxConfig, err := decoder.Decode(decoder.AcquirerService)
	if err != nil {
		log.Fatal(err)
	}
	repo := chooseRepo(trxConfig)
	service := acquirer.NewAcquirerService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/acquirer/payments/{acquirerid}", handler.GetAcquirerPayments)
	r.Get("/merchant/payments/{mid}", handler.GetMerchantPayments)
	r.Get("/acquirer/reports/{acquirerid}", handler.GetReports)
	r.Get("/merchant/transactions/{acquirerid}", handler.GetTransactions)
	r.Get("/rates/{date}", handler.GetExchangeRates)
	r.Get("/logs/{date}", handler.GetLogMessages)
	r.Get("/states/{approved}", handler.GetProcessStates)
	r.Put("/states/{stateid}", handler.PutProcessState)
	// r.Get("/acquirer/transactions/{acquirerid}", handler.GetTransactions)

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

func chooseRepo(rc decoder.Config) acquirer.Repository {

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
