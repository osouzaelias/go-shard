package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-shard/internal/ports"
	"log"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *gin.Engine
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port, server: gin.Default()}
}

func (a Adapter) Run() {
	a.server.GET("/v1/health", a.HealthCheck)
	a.server.GET("/v1/shard", a.Get)

	if err := a.server.Run(fmt.Sprintf(":%d", a.port)); err != nil {
		log.Fatal("error when initializing server", err)
	}
}
