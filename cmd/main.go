package main

import (
	"go-shard/internal/adapters/db"
	"go-shard/internal/adapters/grpc"
	"go-shard/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter("us-east-1", "xx-happy-path")
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, 8081)
	grpcAdapter.Run()
}
