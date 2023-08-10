package server

import "market/market/infrastructure/baseServer"

type Server struct {
	bs baseServer.Server
}

// todo add config
func New(config Config) Server {
	return Server{
		bs: baseServer.New(config.UrlPath, config.Port),
	}
}
