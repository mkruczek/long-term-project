package main

import (
	"context"
	"market/market/application/server"
	"market/market/infrastructure/log"
)

func main() {

	mainCtx := context.Background()

	log.Init("info")

	svr := server.New()

	svr.Init(mainCtx)

	if err := svr.ListenAndServe(mainCtx); err != nil {
		log.Fatalf(mainCtx, "can`t start api server: %s", err)
	}
}
