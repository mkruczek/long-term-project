package server

import "market/market/infrastructure/baseServer"

type Config struct {
	UrlPath string
	Port    string
}

type Server struct {
	bs baseServer.Server
}

func New(config Config) Server {
	return Server{
		bs: baseServer.New(config.UrlPath, config.Port),
	}
}
