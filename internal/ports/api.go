package ports

import (
	"context"
	"go-shard/internal/application/core/domain"
)

type APIPort interface {
	GetCell(ctx context.Context, tenantID, shardID, customerID string) (*domain.Cell, error)
}
