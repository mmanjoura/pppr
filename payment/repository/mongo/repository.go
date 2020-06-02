package mongo

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"log"

	guuid "github.com/google/uuid"
	"github.com/mmanjoura/pppr/payment"
	"github.com/mmanjoura/pppr/transaction"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

const (
	PaymentTask = iota
	StatementTask
	InvoiceTask
	RebateTask
	RefundTask
)

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

// NewMongoRepository ...
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (payment.Repository, error) {
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

// Run ...
func (r *mongoRepository) Run(param *payment.RunParams, colec string, model *payment.PAcquirer) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("transactions")
	col := r.client.Database(r.database).Collection(colec)

	cur, err := collection.Find(ctx, bson.M{"acquirerid": param.AcquirerID})
	defer cur.Close(ctx)

	p := payment.Transaction{}
	task := payment.ProcessState{}

	createdDate := time.Now().Format("2006-01-02")
	createdTime := time.Now().Format("2006.01.02 15:04:05")[11:]
	paymentID := guuid.New()

	//Temp, Remove when deploy

	tempMID := ""
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var trx = transaction.Meta{}

		// decode similar to deserialize process.
		err := cur.Decode(&trx)
		if err != nil {
			log.Fatal(err)
		}
		if tempMID != trx.MHRMerchantIdentifier {
			paymentID = guuid.New()
		}
		tempMID = trx.MHRMerchantIdentifier
		//Look for when MID is different!!

		p.TransactionID = trx.TARTranID
		p.PaymentID = paymentID.String()
		p.CreatedDate = createdDate
		p.CreatedTime = createdTime

		p.AcquirerID = trx.AcquirerID
		p.MerchantID = trx.MHRMerchantIdentifier
		p.TerminalID = trx.TDRTranTerminalID

		// Performing the following formatting according TXF file spec //  00000001225(.01225)
		formattedtom := trx.TDRTranCardholderAmount[:6] + "." + trx.TDRTranCardholderAmount[len(trx.TDRTranCardholderAmount)-5:]
		formattedcha := trx.TDRTranCardholderAmount[:6] + "." + trx.TDRTranCardholderAmount[len(trx.TDRTranCardholderAmount)-5:]
		formattedsfa := trx.TSFSFAmount[:6] + "." + trx.TSFSFAmount[len(trx.TSFSFAmount)-5:]
		formattedifa := trx.TIFIFAmount[:6] + "." + trx.TIFIFAmount[len(trx.TIFIFAmount)-5:]
		formattedafa := trx.TAFAFAmount[:6] + "." + trx.TAFAFAmount[len(trx.TAFAFAmount)-5:]
		originalTrxAmount, err := strconv.ParseFloat(formattedtom, 64)
		if err != nil {
			//return errors.Wrap(err, "repository.payment.Store, Fee Amount is invalid")
			originalTrxAmount = 0
		}
		cha, err := strconv.ParseFloat(formattedcha, 64)
		if err != nil {
			cha = 0
			//return errors.Wrap(err, "repository.payment.Store, Fee Amount is invalid")
		}
		sfa, err := strconv.ParseFloat(formattedsfa, 64)
		if err != nil {
			sfa = 0
			//return errors.Wrap(err, "repository.payment.Store, Fee Amount is invalid")
		}
		ifa, err := strconv.ParseFloat(formattedifa, 64)
		if err != nil {
			ifa = 0
			//return errors.Wrap(err, "repository.payment.Store, Fee Amount is invalid")
		}
		afa, err := strconv.ParseFloat(formattedafa, 64)
		if err != nil {
			afa = 0
			//return errors.Wrap(err, "repository.payment.Store, Fee Amount is invalid")
		}

		p.OriginalTransactionAmount = originalTrxAmount
		p.SettledAmount = cha
		//p.Fee = trx.TAFAFAmount
		// p.IsDcc = trx.TARCardIssCurrency
		p.IsCardPresent = trx.TAFAFCardPresence
		//p.PercentageFee						 = trx.
		// p.PptFee = trx.TAFAFPPT
		p.InterchangeFee = ifa
		p.SchemeFee = sfa
		p.AcquirerFee = afa
		//p.BatchId = trx.BTRBatchSettleCode

		p.CardNumberMasked = trx.TDRTranCardPAN
		p.TransactionDate = trx.TDRTranDate
		p.LocalCurrency = trx.TDRTranMerchantCurrencyCode
		p.ForeignCurrency = trx.TDRTranCardholderCurrency
		//p.ForeignAmount						 //= trx
		p.TransactionType = trx.TDRTranTypeCode
		p.AuthCode = trx.TDRTranAuthCode
		// p.RetailerReference = trx.MHRMerchantClientRef
		// p.PaymentId							 = trx
		// p.IsNetSettled						 = trx
		p.MerchantName = trx.MHRMerchantOutletName
		p.CardSchemeCode = trx.TAFAFCardScheme
		p.CardIssuedCountry = trx.TAFAFRegion
		p.MarginAmount = trx.TDRTranIndMarginvalue
		p.MarginRate = trx.TDRTranIndMarginpercent
		//p.AmpsDateCreated = trx.TDRTraTDRTranDate
		//p.DataSourceCode = trx.TDRDataSourceCode
		p.TransactionCode = trx.TDRTranTypeCode
		// p.AcquirerFeeExtended				 = trx
		// p.SettlementType					 = trx
		// p.CustomerReference = trx.AL1CustomerReferenceNumber
		p.IsDccCurrencyOffered = trx.TAREligibleCurrency
		//p.IsFunded = trx.TDRTranTypeCode
		p.CountryCode = trx.TDRTranCountryCode
		//p.Region = trx.TAFAFRegion
		// p.Channel							 = trx
		// p.IsSecure = trx.TAFAFStatus
		p.ChargeType = trx.TDRTranTypeCode
		// p.ChargePercentage					 = trx
		// p.ChargePpt							 = trx
		// p.Service							 = trx
		// p.MscType							 = trx
		p.IFAmount = trx.TIFIFAmount
		p.IFCurrencyCode = trx.TIFIFCurrencyCode
		// p.IFRate = trx.TIFIFRate
		// p.IFPPT = trx.TIFIFPPT
		p.IFParameterCode = trx.TIFIFParameterCode
		// p.SFAmount = trx.TSFSFAmount
		p.SFCurrencyCode = trx.TSFSFCurrencyCode
		// p.SFRate = trx.TSFSFRate
		// p.SFPPT = trx.TSFSFPPT
		// p.AFAmount = trx.TAFAFAmount
		p.AFCurrencyCode = trx.TAFAFCurrencyCode
		// p.AFRate = trx.TAFAFRate
		// p.AFPPT = trx.TAFAFPPT
		p.AFParameterCode = trx.TAFAFParameterCode
		p.AFCardScheme = trx.TAFAFCardScheme
		p.AFRegion = trx.TAFAFRegion
		p.AFProduct = trx.TAFAFProduct
		p.AFCardPresence = trx.TAFAFCardPresence

		_, err = col.InsertOne(
			ctx,
			p,
		)

	}

	// Now update the Process State
	processTask := r.client.Database(r.database).Collection("state")

	task.ID = guuid.New().String()
	task.ProcessType = PaymentTask
	task.CreatedDate = createdDate
	task.CreatedTime = createdTime
	task.Approved = false
	task.ProcessTypeID = paymentID.String()

	_, err = processTask.InsertOne(
		ctx,
		task,
	)

	if err != nil {
		return errors.Wrap(err, "repository.payment.Store")
	}
	return nil
}
