package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"golang.org/x/exp/slog"

	"gotiny/internal/config"
	"gotiny/internal/core/model"
	"gotiny/internal/data/dto"
	"gotiny/internal/data/util"
)

type DynamodbIpRepository struct {
	client *dynamodb.Client
	config *config.Config
}

func NewDynamodIpRepository(
	client *dynamodb.Client,
	cfg *config.Config,
) *DynamodbIpRepository {
	return &DynamodbIpRepository{
		client: client,
		config: cfg,
	}
}

func (r *DynamodbIpRepository) GetIpDetails(
	ctx context.Context,
	ip string,
) (*model.IpDetails, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Key:       util.FormatKeys(dto.IpPK, dto.IpSKPrefix+ip),
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error getting ip details by id", "err", err)
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	var dto dto.IpDetailsDto
	err = attributevalue.UnmarshalMap(result.Item, &dto)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling ip details", "err", err)
		return nil, err
	}
	d := dto.ToIpDetails()

	return &d, nil
}

func (r *DynamodbIpRepository) SaveIpDetails(ctx context.Context, details model.IpDetails) error {
	dto := dto.IpDetailsDtoFromModel(details)
	avs, err := attributevalue.MarshalMap(dto)
	if err != nil {
		slog.ErrorCtx(ctx, "Error marshalling ip details", "err", err)
		return err
	}
	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.config.IpTableName()),
		Item:      avs,
	})

	if err != nil {
		slog.ErrorCtx(ctx, "Error saving ip details", "err", err)
		return err
	}

	return nil
}
