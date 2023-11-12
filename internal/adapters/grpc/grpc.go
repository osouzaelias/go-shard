package grpc

import (
	"context"
	"fmt"
	pkg "go-shard/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Get(ctx context.Context, request *pkg.GetShardRequest) (*pkg.GetShardResponse, error) {
	shard, err := a.api.GetShard(ctx, request.Tenant, request.CustomerId)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to get. %v ", err)).Err()
	}
	return &pkg.GetShardResponse{
		ShardId: shard.ID,
	}, nil
}
