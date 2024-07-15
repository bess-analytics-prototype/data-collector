package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type BessRepoDb struct {
	dbClient  *dynamodb.Client
	tableName string
}

func (d BessRepoDb) GetData() (Bess, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(d.tableName),
		Limit:     aws.Int32(1),
	}

	result, err := d.dbClient.Scan(context.TODO(), input)
	if err != nil {
		return Bess{}, fmt.Errorf("failed to scan items: %v", err)
	}

	if len(result.Items) == 0 {
		return Bess{}, errors.New("no items found")
	}

	var bess Bess
	err = attributevalue.UnmarshalMap(result.Items[0], &bess)
	if err != nil {
		return Bess{}, fmt.Errorf("failed to unmarshal item: %v", err)
	}

	return bess, nil
}

func (d BessRepoDb) PostData(bess Bess) error {
	// Marshal the Bess struct into a map of DynamoDB attribute values
	item, err := attributevalue.MarshalMap(bess)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %v", err)
	}

	// Define the input for the DynamoDB put item operation
	input := &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item:      item,
	}

	// Perform the put item operation
	_, err = d.dbClient.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to put item: %v", err)
	}
	return nil
}

func NewBessRepoDb(tableName string) (*BessRepoDb, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &BessRepoDb{
		dbClient:  dynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}, nil
}
