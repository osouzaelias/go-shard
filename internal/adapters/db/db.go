package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-shard/internal/application/core/domain"
)

type Shard struct {
	Tenant string `dynamodbav:"tenant"`
	Total  uint8  `dynamodbav:"numberOfShards"`
}

type Adapter struct {
	db        *dynamodb.Client
	tableName string
}

func NewAdapter(region, tableName string) (*Adapter, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	return &Adapter{db: client, tableName: tableName}, nil
}

func (a Adapter) Get(ctx context.Context, tenant string) (*domain.Shard, error) {
	var shardEntity Shard

	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"tenant": &types.AttributeValueMemberS{Value: tenant},
		},
		TableName: aws.String(a.tableName),
	}

	output, err := a.db.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(output.Item, &shardEntity)
	if err != nil {
		return nil, err
	}

	return &domain.Shard{
		Tenant: shardEntity.Tenant,
		Total:  shardEntity.Total,
	}, nil
}
