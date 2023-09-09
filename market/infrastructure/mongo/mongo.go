package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/market/domain"
	"market/market/infrastructure/log"
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

func (c Provider) Insert(ctx context.Context, trade domain.Trade) error {
	coll := c.client.Database(dataBase).Collection(collection)
	_, err := coll.InsertOne(ctx, trade)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) InsertBulk(ctx context.Context, trades []domain.Trade) error {
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

func (c Provider) Get(ctx context.Context, id string) (domain.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	var trade domain.Trade
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&trade)
	if err != nil {
		return domain.Trade{}, err
	}
	return trade, nil
}

func (c Provider) Update(ctx context.Context, trade domain.Trade) error {
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

func (c Provider) List(ctx context.Context) ([]domain.Trade, error) {
	coll := c.client.Database(dataBase).Collection(collection)
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	var trades []domain.Trade
	err = cursor.All(ctx, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil
}
