package main

import (
	"context"
	"market/market/application/statistic/server"
	"market/market/application/statistic/useCase"
	statistic "market/market/domain/statistics_v2"
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

	tradeService := statistic.NewTradesProvider(mongoDB)
	statisticService := statistic.New(tradeService)

	getStsUseCase := useCase.NewStatisticGetter(statisticService, tradeService)

	appServices := server.NewAppServices(getStsUseCase)

	svr := server.New(appServices)

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
