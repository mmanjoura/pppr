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

	"github.com/mmanjoura/pppr/acquirer"
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

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (acquirer.Repository, error) {
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

func (r *mongoRepository) GetAcquirerPayments(acquirerId string) ([]acquirer.Payment, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("payments")

	payments := []acquirer.Payment{}
	findOptions := options.Find()
	// findOptions.SetLimit(200)

	cur, err := collection.Find(ctx, bson.M{"acquirerid": acquirerId}, findOptions)
	defer cur.Close(ctx)

	var p = acquirer.Payment{}
	for cur.Next(ctx) {

		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		payments = append(payments, p)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return payments, nil
}

func (r *mongoRepository) GetMerchantPayments(MID string) ([]acquirer.Payment, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("payments")

	payments := []acquirer.Payment{}
	findOptions := options.Find()
	// findOptions.SetLimit(200)

	cur, err := collection.Find(ctx, bson.M{"merchantid": MID}, findOptions)
	defer cur.Close(ctx)

	var p = acquirer.Payment{}
	for cur.Next(ctx) {

		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		payments = append(payments, p)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return payments, nil
}

func (r *mongoRepository) GetReports(acquirerId string) ([]acquirer.Report, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("reports")

	reports := []acquirer.Report{}
	findOptions := options.Find()
	// findOptions.SetLimit(100)

	cur, err := collection.Find(ctx, bson.M{"acquirerid": acquirerId}, findOptions)
	defer cur.Close(ctx)

	var rep = acquirer.Report{}
	for cur.Next(ctx) {

		err := cur.Decode(&rep)
		if err != nil {
			log.Fatal(err)
		}
		reports = append(reports, rep)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return reports, nil
}

func (r *mongoRepository) GetTransactions(acquirerId string) ([]acquirer.Transaction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("transactions")

	transactions := []acquirer.Transaction{}
	findOptions := options.Find()
	// findOptions.SetLimit(100)

	cur, err := collection.Find(ctx, bson.M{"acquirerid": acquirerId}, findOptions)
	defer cur.Close(ctx)

	var t = acquirer.Transaction{}
	for cur.Next(ctx) {

		err := cur.Decode(&t)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, t)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return transactions, nil
}

func (r *mongoRepository) GetExchangeRates(date string) (acquirer.ExchangeRate, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("rates")

	findOptions := options.Find()
	// findOptions.SetLimit(100)

	cur, err := collection.Find(ctx, bson.M{"createddate": date}, findOptions)
	defer cur.Close(ctx)

	var exr = acquirer.ExchangeRate{}
	for cur.Next(ctx) {

		err := cur.Decode(&exr)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		return acquirer.ExchangeRate{}, errors.Wrap(err, "repository.acquirer.Store")
	}
	return exr, nil
}

func (r *mongoRepository) GetLogMessages(date string) ([]acquirer.Message, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("logs")

	logMessages := []acquirer.Message{}
	findOptions := options.Find()
	// findOptions.SetLimit(100)

	cur, err := collection.Find(ctx, bson.M{"createddate": date}, findOptions)
	defer cur.Close(ctx)

	var lg = acquirer.Message{}
	for cur.Next(ctx) {

		err := cur.Decode(&lg)
		if err != nil {
			log.Fatal(err)
		}
		logMessages = append(logMessages, lg)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return logMessages, nil
}

func (r *mongoRepository) GetProcessStates(approved bool) ([]acquirer.ProcessState, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("state")

	processStates := []acquirer.ProcessState{}
	findOptions := options.Find()
	// findOptions.SetLimit(100)

	cur, err := collection.Find(ctx, bson.M{"approved": approved}, findOptions)
	defer cur.Close(ctx)

	var ps = acquirer.ProcessState{}
	for cur.Next(ctx) {

		err := cur.Decode(&ps)
		if err != nil {
			log.Fatal(err)
		}
		processStates = append(processStates, ps)
	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.acquirer.Store")
	}
	return processStates, nil
}

func (r *mongoRepository) PutProcessState(processID string, ps acquirer.ProcessState) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("state")

	filter := bson.M{"processtypeid": bson.M{"$eq": processID}}
	update := bson.M{"$set": bson.M{"approved": "true"}}

	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return err

	}

	return nil
}
