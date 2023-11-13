package api

import (
	"context"
	"fmt"
	"github.com/cespare/xxhash/v2"
	"go-shard/internal/application/core/domain"
	"go-shard/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) GetShard(_ context.Context, tenant, customerID string) (*domain.Shard, error) {
	shards := make([]string, 0)
	for i := uint8(0); i < 10; i++ {
		shards = append(shards, fmt.Sprintf("Shard%d", i))
	}

	rendezvous := domain.NewRendezvous(shards, xxhash.Sum64String)

	return &domain.Shard{
		ID:     rendezvous.Lookup(customerID),
		Tenant: tenant,
		Total:  10,
	}, nil
}

//func (a Application) GetShard(ctx context.Context, tenant, customerID string) (*domain.Shard, error) {
//	shard, _ := a.db.Get(ctx, tenant)
//	shards := make([]string, 0)
//	for i := uint8(0); i < shard.Total; i++ {
//		shards = append(shards, fmt.Sprintf("Shard%d", i))
//	}
//
//	rendezvous := domain.NewRendezvous(shards, xxhash.Sum64String)
//	shard.ID = rendezvous.Lookup(customerID)
//
//	return shard, nil
//}
