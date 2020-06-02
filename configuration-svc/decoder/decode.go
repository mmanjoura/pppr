package decoder

import (
	"encoding/json"
	"net/http"
)

var config = Config{}

// Decode ..
func Decode(servicename string) (Config, error) {

	APIURL := "http://localhost:8081/read/" + servicename
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		return Config{}, err
	}
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return Config{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&config)
	return config, nil
}

// Config struct need to stay in Sync with Yaml file, a little pain but offers
// a good value, if gives us strong typing while accessing config vaules from service
type Config struct {
	APIGateway     string `json:"apiGateway,omitempty"`
	MmsEndPoint    string `json:"mmsendpoint,omitempty"`
	MmsBaseURL     string `json:"mmsbaseurl,omitempty"`
	HTTPServer     string `json:"httpServer,omitempty"`
	GRPCServer     string `json:"gRPCServer,omitempty"`
	HTTPPort       int    `json:"httpPort,omitempty"`
	GRPCPort       int    `json:"grpcPort,omitempty"`
	SourcePath     string `json:"sourcePath,omitempty"`
	ArchivePath    string `json:"archivePath,omitempty"`
	DropFolder     string `json:"dropfolder,omitempty"`
	ReportImages   string `json:"reportimages,omitempty"`
	ScanInterval   int    `json:"scanInterval,omitepty"`
	DbHost         string `json:"dbHost,omitempty"`
	DbName         string `son:"dbName,omitempty"`
	CollectionName string `json:"collectionName,omitempty"`

	Mongo struct {
		DbHost       string `json:"dbHost,omitempty"`
		Port         string `json:"port,omitempty"`
		Mongourl     string `json:"mongoUrl,omitempty"`
		MongoTimeout int    `json:"mongoTimeout,omitempty"`
	} `json:"mongo"`
	Redis struct {
		Host     string `json:"host,omitempty"`
		Port     int    `json:"port,omitempty"`
		Redisurl string `json:"redisUrl,omitempty"`
	} `json:"redis"`
	BoltDb struct {
		DBHost string `json:"dbHost,omitempty"`
		Port   string `json:"port,omitempty"`
		Path   string `jon:"path,omitempty"`
	} `json:"boltDb"`
}

// Message structure r Model
type Message struct {
	CreatedAt     string `json:"created_at" bson:"creaed_at"`
	Level         string `json:"level"   bson:"level"`
	ServiceName   string `json:"service_name"   bson:"service_name"`
	CallingMethod string `json:"calling_method"   bsn:"calling_method"`
	Host          string `json:"host"   bson:"host"`
	Body          string `json:"body"   bson:"body"`
	Latency       string `json:"latency"   bson:"latency"`
}

//ConfigBase ...
const ConfigBase = "base"

/************* base of the yaml file, wil always get returned ***************

base:
  apiGatway: http://service.api-gateway
  redis:
    dbHost: reis
    port: 6379
    rediurl: redis://localhost
  mongo:
    dbHost: mono
    port: 27017
    mongourl: mongod://localhost
    monotimeout: 30
  Bolt:
    dbHost: bot
    port: 000
   path: "

******************************************************************************/
// PaymentService ...
const LoggingService = "service.logging"

/********************* Payment Processor Service configuration items ***********

service.logging:
  httpServer: localhost
  gRPCServer: localhost
  httpport: 8000
  gRPCPort: 50050
  dbHost: mongo
  dbName: pppr_services_Db
  collectionName: logs

******************************************************************************/

// TransactionService ...
const TransactionService = "service.transaction"

/********************* Transaction Service configuration items *****************

service.transaction:
  httpServer: localhost
  gRPCServer: localhost
  httpport: 8010
  gRPCPort: 50060
  sourcePath: ./responsefiles/dataFiles/
  archivePath: ./responsefiles/archive/
  scanInterval: 30
  dbHost: mongo
  dbName: pppr_services_Db
  collectionName: transactions

******************************************************************************/

// PaymentService ...
const PaymentService = "service.payment"

/********************* Payment Service config items************** **************

service.payment
 httpServer: localhost
  gRPCServer: localhost
  httpport: 8020
  gRPCPort: 50070
  dbHost: mongo
  dbName: pppr_services_Db
  collectionName: payments

******************************************************************************/

// ReportService ...
const ReportService = "service.report"

/********************* Payment Service config items************** **************

service.report:
  httpServer: localhost
  gRPCServer: localhost
  httpport: 8030
  gRPCPort: 50080
  dbHost: mongo
  dbName: pppr_services_Db
  collectionName: reports
  reportimages: ../../../images/

******************************************************************************/
// AcquirerService ...
const AcquirerService = "service.acquirer"

/******************** Payment Service config items************** **************

service.acquirer:
  httpServer: localhost
  gRPCServer: localhost
  httpport: 8040
  gRPCPort: 50090
  dbHost: mongo
  dbName: pppr_services_Db
  collectionName: ""

*******************************************************************************/
// DrgService ...
const DrgService = "service.drg"

/******************** Payment Service config items************** **************

service.drg:
  httpServer: localhost
  gRPCServer: localhost
  httpport: 8050
  gRPCPort: 50100
  dbHost: mongo
  scanInterval: 30
  dbName: pppr_services_Db
  collectionName: "exchangeRates"

*******************************************************************************/
