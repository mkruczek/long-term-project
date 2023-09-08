package main

import (
	"context"
	"market/market/application/server"
	"market/market/domain"
	"market/market/domain/dataProviders/xtb"
	"market/market/infrastructure/log"
	"market/market/infrastructure/mongo"
)

func main() {

	mainCtx := context.Background()
	log.Init("info")

	//todo add config
	mongoDB, err := mongo.New(mainCtx, "market", "localhost", "27017", "root", "secret")
	if err != nil {
		log.Fatalf(mainCtx, "can`t create mongo provider: %s", err)
	}

	xtbProv := domain.NewProvider[*xtb.CSV](mongoDB)
	tradeSrv := domain.NewService(mongoDB)

	appServices := server.NewAppServices(tradeSrv, xtbProv)

	svr := server.New(appServices)

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
