package ports

import (
	"context"
	"go-shard/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (*[]domain.Cell, error)
}
