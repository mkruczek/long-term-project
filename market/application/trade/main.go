package main

import (
	"context"
	"market/market/application/trade/server"
	"market/market/domain/dataProviders/xtb"
	"market/market/domain/trade"
	"market/market/libs/config"
	"market/market/libs/log"
	"market/market/libs/mongo"
)

func main() {

	mainCtx := context.Background()
	log.Init("info")

	mgoCfg := config.GetTrade().DataBase
	mongoDB, err := mongo.New(mainCtx, mgoCfg.DBName, mgoCfg.Host, mgoCfg.Port, mgoCfg.User, mgoCfg.Password)
	if err != nil {
		log.Fatalf(mainCtx, "can`t create mongo provider: %s", err)
	}

	xtbProv := trade.NewProvider[*xtb.CSV](mongoDB)
	tradeSrv := trade.NewService(mongoDB)

	appServices := server.NewAppServices(tradeSrv, xtbProv)

	svr := server.New(appServices)

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
