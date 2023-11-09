package main

import (
	"context"
	"market/market/application/server"
	"market/market/domain"
	"market/market/domain/dataProviders/xtb"
	"market/market/libs/config"
	"market/market/libs/log"
	"market/market/libs/mongo"
)

func main() {

	mainCtx := context.Background()
	log.Init("info")

	mgoCfg := config.GetMarket().Mongo
	mongoDB, err := mongo.New(mainCtx, mgoCfg.DBName, mgoCfg.Host, mgoCfg.Port, mgoCfg.User, mgoCfg.Password)
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
