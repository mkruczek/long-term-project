package main

import (
	"context"
	"market/market/application/server"
	"market/market/domain/tradeProvider/xtb"
	"market/market/infrastructure/log"
	"market/market/infrastructure/mongo"
)

func main() {

	mainCtx := context.Background()
	log.Init("info")

	mongoDB, err := mongo.New(mainCtx, "market", "localhost", "27017", "root", "secret")
	if err != nil {
		log.Fatalf(mainCtx, "can`t create mongo provider: %s", err)
	}

	xtbProv := xtb.NewProvider(mongoDB)

	appServices := server.NewAppServices(xtbProv)

	svr := server.New(appServices)

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
