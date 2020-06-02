package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"log"

	guuid "github.com/google/uuid"
	"github.com/mmanjoura/pppr/payment"
	"github.com/mmanjoura/pppr/report"
)

const (
	PaymentTask = iota
	StatementTask
	InvoiceTask
	RebateTask
	RefundTask
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (report.Repository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}

func (r *mongoRepository) GetPayments(acquirerId, collc string) ([]report.Payment, map[string]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("payments")
	//col := r.client.Database(r.database).Collection(collc)

	//merchants := []report.Merchant{}
	payments := []report.Payment{}

	cur, err := collection.Find(ctx, bson.M{"acquirerid": acquirerId})
	defer cur.Close(ctx)

	var trx = payment.Transaction{}

	var p = report.Payment{}

	p.CreatedDate = time.Now().Format("2006-01-02")
	p.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[11:]

	tempMID := ""
	mids := make(map[string]string)
	for cur.Next(ctx) {

		err := cur.Decode(&trx)
		if err != nil {
			log.Fatal(err)
		}

		if tempMID != trx.MerchantID {
			mids[trx.MerchantID] = trx.MerchantID
		}

		p.TransactionID = trx.TransactionID
		p.AcquirerID = trx.AcquirerID
		p.MerchantID = trx.MerchantID
		p.TerminalID = trx.TerminalID
		p.PaymentID = trx.PaymentID

		p.OriginalTransactionAmount = trx.OriginalTransactionAmount
		p.SettledAmount = trx.OriginalTransactionAmount
		//p.Fee = trx.TAFAFAmount
		// p.IsDcc = trx.TARCardIssCurrency
		p.IsCardPresent = trx.IsCardPresent
		//p.PercentageFee						 = trx.
		// p.PptFee = trx.TAFAFPPT
		p.InterchangeFee = trx.InterchangeFee
		p.SchemeFee = trx.SchemeFee
		p.AcquirerFee = trx.AcquirerFee
		//p.BatchId = trx.BTRBatchSettleCode

		p.CardNumberMasked = trx.CardNumberMasked
		p.TransactionDate = trx.TransactionDate
		p.LocalCurrency = trx.LocalCurrency
		p.ForeignCurrency = trx.ForeignCurrency
		//p.ForeignAmount						 //= trx
		p.TransactionType = trx.TransactionType
		p.AuthCode = trx.AuthCode
		// p.RetailerReference = trx.MHRMerchantClientRef
		// p.PaymentId							 = trx
		// p.IsNetSettled						 = trx
		p.MerchantName = trx.MerchantName
		p.CardSchemeCode = trx.CardSchemeCode
		p.CardIssuedCountry = trx.CardIssuedCountry
		p.MarginAmount = trx.MarginAmount
		p.MarginRate = trx.MarginRate
		//p.AmpsDateCreated = trx.TDRTraTDRTranDate
		//p.DataSourceCode = trx.TDRDataSourceCode
		p.TransactionCode = trx.TransactionCode
		// p.AcquirerFeeExtended				 = trx
		// p.SettlementType					 = trx
		// p.CustomerReference = trx.AL1CustomerReferenceNumber
		p.IsDccCurrencyOffered = trx.IsDccCurrencyOffered
		//p.IsFunded = trx.TDRTranTypeCode
		p.CountryCode = trx.CountryCode
		//p.Region = trx.TAFAFRegion
		// p.Channel							 = trx
		// p.IsSecure = trx.TAFAFStatus
		p.ChargeType = trx.ChargeType
		// p.ChargePercentage					 = trx
		// p.ChargePpt							 = trx
		// p.Service							 = trx
		// p.MscType							 = trx
		p.IFAmount = trx.IFAmount
		p.IFCurrencyCode = trx.IFCurrencyCode
		// p.IFRate = trx.TIFIFRate
		// p.IFPPT = trx.TIFIFPPT
		p.IFParameterCode = trx.IFParameterCode
		// p.SFAmount = trx.TSFSFAmount
		p.SFCurrencyCode = trx.SFCurrencyCode
		// p.SFRate = trx.TSFSFRate
		// p.SFPPT = trx.TSFSFPPT
		// p.AFAmount = trx.TAFAFAmount
		p.AFCurrencyCode = trx.AFCurrencyCode
		// p.AFRate = trx.TAFAFRate
		// p.AFPPT = trx.TAFAFPPT
		p.AFParameterCode = trx.AFParameterCode
		p.AFCardScheme = trx.AFCardScheme
		p.AFRegion = trx.AFRegion
		p.AFProduct = trx.AFProduct
		p.AFCardPresence = trx.AFCardPresence
		payments = append(payments, p)

	}

	if err != nil {
		return nil, nil, errors.Wrap(err, "repository.repot.Store")
	}

	//rm.Payments = payments
	//merchants = append(merchants, rm)
	return payments, mids, nil

}

func (r *mongoRepository) GenerateReports(reports []report.Report, acquirerId, collec string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(collec)
	reportID := guuid.New()
	createdDate := time.Now().Format("2006-01-02")
	createdTime := time.Now().Format("2006.01.02 15:04:05")[11:]

	task := report.ProcessState{}
	tempMid := ""
	for _, reprt := range reports {

		if tempMid != reprt.MerchantID {
			reprt.ReportID = reportID.String()
			_, err := collection.InsertOne(
				ctx,
				reprt,
			)
			if err != nil {
				return errors.Wrap(err, "repository.Transaction.Save")
			}

		}

		tempMid = reprt.MerchantID

	}

	processTask := r.client.Database(r.database).Collection("state")

	task.ID = guuid.New().String()
	task.ProcessType = StatementTask
	task.CreatedDate = createdDate
	task.CreatedTime = createdTime
	task.Approved = false
	task.ProcessTypeID = reportID.String()

	_, err := processTask.InsertOne(
		ctx,
		task,
	)
	if err != nil {
		return errors.Wrap(err, "repository.Report.Store")
	}

	return nil
}
