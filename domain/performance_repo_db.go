package domain

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PerformanceDataRepoDb struct {
	dbClient  *dynamodb.Client
	tableName string
}

func (d PerformanceDataRepoDb) PostData(performanceData []PerformanceData) error {
	writeRequests := []types.WriteRequest{}

	for _, data := range performanceData {
		item, err := attributevalue.MarshalMap(data)
		if err != nil {
			return fmt.Errorf("failed to marshal item: %v", err)
		}

		writeRequests = append(writeRequests, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		})
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			d.tableName: writeRequests,
		},
	}

	_, err := d.dbClient.BatchWriteItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to batch write items: %v", err)
	}

	return nil
}

func NewPerformanceDataRepoDb(tableName string) (*PerformanceDataRepoDb, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &PerformanceDataRepoDb{
		dbClient:  dynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}, nil
}
