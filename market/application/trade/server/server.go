package server

import (
	"context"
	"market/market/libs/baseServer"
	"market/market/libs/config"
	"market/market/libs/healthReady"
	"market/market/libs/log"
	"market/market/libs/mongo"
)

type Server struct {
	baseServer.Server
	services appServices
}

func New(services appServices) *Server {
	return &Server{
		Server:   baseServer.New(config.GetTrade().Http.URLPath, config.GetTrade().Http.Port),
		services: services,
	}
}

func (svr *Server) Init(ctx context.Context) {

	db, err := mongo.New(ctx, config.GetTrade().DataBase.DBName, config.GetTrade().DataBase.Host, config.GetTrade().DataBase.Port, config.GetTrade().DataBase.User, config.GetTrade().DataBase.Password)
	if err != nil {
		log.Fatalf(ctx, "can`t connect to mongo: %s", err)
	}

	svr.Routes()

	svr.RegisterHealthReady(healthReady.New(
		healthReady.Observer{
			ServiceContextName: "database",
			Service:            db,
		}))
}
