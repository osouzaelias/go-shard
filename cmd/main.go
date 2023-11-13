package main

import (
	"go-shard/config"
	"go-shard/internal/adapters/db"
	"go-shard/internal/adapters/rest"
	"go-shard/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetRegion(), config.GetTableName())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	//application := api.NewApplication(dbAdapter)
	//grpcAdapter := grpc.NewAdapter(application, config.GetPort())
	//grpcAdapter.Run()

	application := api.NewApplication(dbAdapter)
	restAdapter := rest.NewAdapter(application, config.GetPort())
	restAdapter.Run()
}
