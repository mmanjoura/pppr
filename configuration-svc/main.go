package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jakewright/muxinator"
	"github.com/mmanjoura/pppr/configuration-svc/controller"
	"github.com/mmanjoura/pppr/configuration-svc/domain"
	"github.com/mmanjoura/pppr/configuration-svc/service"
)

// http://localhost:8081/read/base
// http://localhost:8081/read/service.logging
// http://localhost:8081/read/service.transaction
// http://localhost:8081/read/service.payment
// http://localhost:8081/read/service.report
// http://localhost:8081/read/service.acquirer
// http://localhost:8081/read/service.drg
var (
	hotName string
)

func main() {
	hotName, err := os.Hostname()
	if err != nil {
		log.Fatal("Failed to start Configuration Service, Cannot get host name")
	}

	config := domain.Config{}

	configService := service.ConfigService{
		Config:   &config,
		Location: "config.yaml",
	}

	go configService.Watch(time.Second * 3)

	c := controller.Controller{
		Config: &config,
	}

	//fmt.Println("Running Config Server on Port: 8081")
	fmt.Printf("Running Configuration Servie on %s Port: 8081", hotName)

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", http.HandlerFunc(c.ReadConfig))
	log.Fatal(router.ListenAndServe(":8081"))

}
