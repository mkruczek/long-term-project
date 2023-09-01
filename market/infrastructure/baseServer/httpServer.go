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
	*gin.Engine
	Port        string
	UrlPath     string
	HealthReady healthReady
}

func New(urlPath, port string) Server {
	return Server{
		Engine:  gin.New(),
		UrlPath: urlPath,
		Port:    port,
	}
}

func (bs *Server) ListenAndServe(ctx context.Context) error {
	return bs.Run(fmt.Sprintf(":%s", bs.Port))
}

func (bs *Server) RegisterHealthReady(hr healthReady) {
	bs.HealthReady = hr

	bs.GET(fmt.Sprintf("%s/health", bs.UrlPath), health(bs.HealthReady))
	bs.GET(fmt.Sprintf("%s/ready", bs.UrlPath), ready(bs.HealthReady))
}
