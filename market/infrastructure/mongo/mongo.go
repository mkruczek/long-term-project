package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/market/domain"
	"market/market/infrastructure/log"
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
	coll := c.client.Database("market").Collection("trades")
	_, err := coll.InsertOne(ctx, trade)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) Read(ctx context.Context, id string) (domain.Trade, error) {
	coll := c.client.Database("market").Collection("trades")
	var trade domain.Trade
	err := coll.FindOne(ctx, domain.Trade{ID: id}).Decode(&trade)
	if err != nil {
		return domain.Trade{}, err
	}
	return trade, nil
}

func (c Provider) Update(ctx context.Context, trade domain.Trade) error {
	coll := c.client.Database("market").Collection("trades")
	_, err := coll.UpdateOne(ctx, domain.Trade{ID: trade.ID}, trade)
	if err != nil {
		return err
	}
	return nil
}

func (c Provider) Delete(ctx context.Context, id string) error {
	coll := c.client.Database("market").Collection("trades")
	_, err := coll.DeleteOne(ctx, domain.Trade{ID: id})
	if err != nil {
		return err
	}
	return nil
}
