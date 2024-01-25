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

func (a Adapter) Get(ctx context.Context, id string) (*[]domain.Cell, error) {
	var cells = make([]domain.Cell, 0)

	partitionKey := expression.Key("id").Equal(expression.Value(id))
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

	err = attributevalue.UnmarshalListOfMaps(out.Items, &cells)
	if err != nil {
		return nil, err
	}

	return &cells, nil
}
