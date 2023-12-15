package api

import (
	"context"
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

func (a Application) GetShard(ctx context.Context, tenantID, customerID string) (*domain.Shard, error) {
	shardNames := make([]string, 0)

	shards, _ := a.db.Get(ctx, tenantID)
	for _, s := range *shards {
		shardNames = append(shardNames, s.ShardID)
	}

	rendezvous := domain.NewRendezvous(shardNames, xxhash.Sum64String)

	shardID := rendezvous.Lookup(customerID)
	for _, s := range *shards {
		if s.ShardID == shardID {
			return &domain.Shard{
				ShardID:  shardID,
				TenantID: s.TenantID,
				Address:  s.Address,
			}, nil
		}
	}

	return &domain.Shard{}, nil
}
