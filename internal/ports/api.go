package ports

import (
	"context"
	"go-shard/internal/application/core/domain"
)

type APIPort interface {
	GetShard(ctx context.Context, tenant, customerID string) (*domain.Shard, error)
}
