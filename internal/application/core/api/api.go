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

var cells *[]domain.Cell
var rz *domain.Rendezvous

func (a Application) GetCell(ctx context.Context, tenantID, shardID, customerID string) (*domain.Cell, error) {
	cellNames := make([]string, 0)

	if cells == nil {
		cells, _ = a.db.Get(ctx, domain.CreateCellID(tenantID, shardID))
		for _, c := range *cells {
			cellNames = append(cellNames, c.Name)
		}
		rz = domain.NewRendezvous(cellNames, xxhash.Sum64String)
	}

	name := rz.Lookup(customerID)

	for _, c := range *cells {
		if c.Name == name {
			return &c, nil
		}
	}

	return &domain.Cell{}, nil
}
