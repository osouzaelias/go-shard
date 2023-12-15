package rest

import (
	"fmt"
	"go-shard/internal/ports"
	"log"
	"net/http"
)

type Adapter struct {
	api  ports.APIPort
	port int
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	http.HandleFunc("/health", a.HealthCheck)
	http.HandleFunc("/", a.Proxy)

	println("Proxy server running on http://localhost:", a.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), nil))
}
