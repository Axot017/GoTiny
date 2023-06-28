package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func newAwsConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("eu-central-1"))
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func newDynamobdClient(cfg *aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(*cfg)
}
