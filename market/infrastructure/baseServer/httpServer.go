package baseServer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type healthReady interface {
	Health(ctx context.Context) bool
	Ready(ctx context.Context) map[string]error
}

type Server struct {
	svr         *gin.Engine
	Port        string
	UrlPath     string
	HealthReady healthReady
}

func New(urlPath, port string) Server {
	return Server{
		svr:     gin.New(),
		UrlPath: urlPath,
		Port:    port,
	}
}

func (bs *Server) ListenAndServe(ctx context.Context) error {
	return bs.svr.Run(fmt.Sprintf(":%s", bs.Port))
}

func (bs *Server) RegisterHealthReady(hr healthReady) {
	bs.HealthReady = hr

	bs.svr.GET(fmt.Sprintf("%s/health", bs.UrlPath), health(bs.HealthReady))
	bs.svr.GET(fmt.Sprintf("%s/ready", bs.UrlPath), ready(bs.HealthReady))
}
