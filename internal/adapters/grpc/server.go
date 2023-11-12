package grpc

import (
	"fmt"
	"go-shard/internal/ports"
	pkg "go-shard/pkg/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	pkg.UnimplementedShardServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	a.server = grpcServer
	pkg.RegisterShardServer(grpcServer, a)
	//if config.GetEnv() == "development" {
	//	reflection.Register(grpcServer)
	//}

	log.Printf("starting balance-shard service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
