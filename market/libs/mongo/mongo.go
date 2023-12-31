package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/market/domain/trade"
	"market/market/libs/log"
	"time"
)

const (
	dataBase   = "market"
	collection = "trades"
)

type Provider struct {
	client *mongo.Client
}

func New(ctx context.Context, dbName, host, port, username, password string) (*Provider, error) {
	connectionUrl := fmt.Sprintf("mongodb://%s:%s", host, port)
	credential := options.Credential{Username: username, Password: password}
	mongoOptions := options.Client().ApplyURI(connectionUrl).SetAppName(dbName).SetAuth(credential)

	mgoClient, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Errorf(ctx, "Unable to create mgoClient %v", err)
		return nil, err
	}

	provider := &Provider{client: mgoClient}

	if err = provider.Ping(ctx); err != nil {
		log.Errorf(ctx, "Unable to verify connection to mongo %v", err)
		return nil, err
	}

	return provider, nil
}

func (c Provider) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func (c Provider) Insert(ctx context.Context, trade trade.Trade) error {
	coll := c.client.Database(dataBase).Collection(collection)
	_, err := coll.InsertOne(ctx, trade)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) BulkInsert(ctx context.Context, trades []trade.Trade) error {
	coll := c.client.Database(dataBase).Collection(collection)
	docs := make([]interface{}, len(trades))
	for i, trade := range trades {
		docs[i] = trade
	}
	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) Get(ctx context.Context, id string) (trade.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	var t trade.Trade
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&t)
	if err != nil {
		return trade.Trade{}, err
	}
	return t, nil
}

func (c Provider) Update(ctx context.Context, trade trade.Trade) error {
	coll := c.client.Database(dataBase).Collection(collection)
	_, err := coll.UpdateOne(ctx, bson.D{{"_id", trade.ID}}, trade)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) Delete(ctx context.Context, id string) error {
	coll := c.client.Database(dataBase).Collection(collection)
	_, err := coll.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) GetAll(ctx context.Context) ([]trade.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var trades []trade.Trade
	err = cursor.All(ctx, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

func (c Provider) GetRange(ctx context.Context, startTime, endTime time.Time) ([]trade.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	cursor, err := coll.Find(ctx,
		bson.D{
			{"openTime", bson.D{{"$gt", startTime}, {"$lt", endTime}}},
			{"closeTime", bson.D{{"$gt", startTime}, {"$lt", endTime}}},
		},
	)
	if err != nil {
		return nil, err
	}
	var trades []trade.Trade
	err = cursor.All(ctx, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}

func (c Provider) GetRangeAndSymbol(ctx context.Context, startTime, endTime time.Time, symbol string) ([]trade.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	cursor, err := coll.Find(ctx,
		bson.D{
			{"openTime", bson.D{{"$gt", startTime}, {"$lt", endTime}}},
			{"closeTime", bson.D{{"$gt", startTime}, {"$lt", endTime}}},
			{"symbol", bson.D{{"$eq", symbol}}},
		},
	)
	if err != nil {
		return nil, err
	}
	var trades []trade.Trade
	err = cursor.All(ctx, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}
