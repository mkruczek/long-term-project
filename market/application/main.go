package main

import (
	"context"
	"market/market/application/server"
	"market/market/infrastructure/log"
	"market/market/infrastructure/tradeaProvider/xtb"
)

func main() {

	mainCtx := context.Background()
	log.Init("info")

	xtbProv := xtb.NewProvider()

	appServices := server.NewAppServices(xtbProv)

	svr := server.New(appServices)

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
