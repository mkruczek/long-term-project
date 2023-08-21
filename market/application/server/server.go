package server

import (
	"context"
	"market/market/infrastructure/baseServer"
	"market/market/infrastructure/config"
	"market/market/infrastructure/healthReady"
	"market/market/infrastructure/log"
	"market/market/infrastructure/mongo"
)

type Server struct {
	baseServer.Server
	services appServices
}

func New(services appServices) *Server {
	return &Server{
		Server:   baseServer.New(config.GetMarket().Http.URLPath, config.GetMarket().Http.Port),
		services: services,
	}
}

func (svr *Server) Init(ctx context.Context) {

	db, err := mongo.New(ctx, config.GetMarket().Mongo.DBName, config.GetMarket().Mongo.Host, config.GetMarket().Mongo.Port, config.GetMarket().Mongo.User, config.GetMarket().Mongo.Password)
	if err != nil {
		log.Fatalf(ctx, "can`t connect to mongo: %s", err)
	}

	svr.RegisterHealthReady(healthReady.New(
		healthReady.Observer{
			ServiceContextName: "database",
			Service:            db,
		}))
}

func (svr *Server) Routes() {
	//todo route with upload file
}
