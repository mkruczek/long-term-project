package main

import (
	"context"
	"market/market/infrastructure/baseServer"
	"market/market/infrastructure/healthReady"
	"market/market/infrastructure/log"
	"market/market/infrastructure/mongo"
)

func main() {

	ctx := context.Background()

	log.Init("info")

	//todo add config
	bs := baseServer.New("/api/market/", "8090")

	db, err := mongo.New(ctx, "market", "172.200.0.10", "27017", "root", "secret")
	if err != nil {
		log.Fatalf(ctx, "can`t connect to mongo: %s", err)
	}

	bs.RegisterHealthReady(healthReady.New(
		healthReady.Observer{
			ServiceContextName: "database",
			Service:            db,
		}))

	if err := bs.ListenAndServe(ctx); err != nil {
		log.Fatalf(ctx, "can`t start api server: %s", err)
	}
}
