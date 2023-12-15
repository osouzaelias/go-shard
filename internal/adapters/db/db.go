package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

func (a Adapter) Get(ctx context.Context, tenantID string) (*[]domain.Shard, error) {
	var shards = make([]domain.Shard, 0)

	partitionKey := expression.Key("tenantID").Equal(expression.Value(tenantID))
	expr, err := expression.NewBuilder().WithKeyCondition(partitionKey).Build()

	input := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(a.tableName),
	}

	out, err := a.db.Query(ctx, input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &shards)
	if err != nil {
		return nil, err
	}

	return &shards, nil
}
