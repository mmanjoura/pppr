package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/mmanjoura/pppr/transaction"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (transaction.Repository, error) {
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

func (r *mongoRepository) Save(trx *transaction.Meta, collec string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(collec)
	dateTime := make(map[string]string)
	dateTime[trx.CreatedTime] = trx.CreatedTime
	_, err := collection.InsertOne(
		ctx,
		trx,
	)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Transaction.Save")
	}
	return dateTime, nil
}

func (r *mongoRepository) Get(date, time string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("transactions")
	acquirers := make([]string, 0)

	//query := `{"$eq":"last value"}`
	// cur, err := collection.Find( { "$and": [ { "createddate" : { "$eq": "2020-05-09" } }, { "createdtime" : { "$eq": "13:03:33" } } ] } )

	cur, err := collection.Find(ctx, bson.D{
		{"createddate", bson.D{
			{"$eq", date},
		}},
		{"createdtime", bson.D{
			{"$eq", time},
		}},
	}) //.Distinct("UHRAcquirerID", &acquirers)
	tempAcquirer := ""
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var trx = transaction.Meta{}

		// decode similar to deserialize process.
		err := cur.Decode(&trx)
		if err != nil {
			return nil, errors.Wrap(err, "repository.Transaction.GetAcquirers")
		}

		if tempAcquirer != trx.AcquirerID {
			acquirers = append(acquirers, trx.AcquirerID)
			fmt.Println(trx.AcquirerID)
		}
		tempAcquirer = trx.AcquirerID

	}

	if err != nil {
		return nil, errors.Wrap(err, "repository.Transaction.GetAcquirers")
	}
	return acquirers, nil
}
