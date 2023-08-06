package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
