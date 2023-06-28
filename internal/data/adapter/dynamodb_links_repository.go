package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/exp/slog"

	"gotiny/internal/core/model"
)

const (
	linksGlobalDataPK  = "links_global"
	linksGlobalIndexSK = "index"
)

type DynamodbConfig interface {
	LinksTableName() string
}

type DynamodbLinksRepository struct {
	client *dynamodb.Client
	config DynamodbConfig
}

func NewDynamodbLinksRepository(
	client *dynamodb.Client,
	cfg DynamodbConfig,
) *DynamodbLinksRepository {
	return &DynamodbLinksRepository{client: client, config: cfg}
}

func (r *DynamodbLinksRepository) GetNextLinkIndex(ctx context.Context) (uint, error) {
	result, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Key: map[string]types.AttributeValue{
			"PK": types.AttributeValueMemberS{
				Value: linksGlobalDataPK,
			},
			"SK": types.AttributeValueMemberS{
				Value: linksGlobalIndexSK,
			},
		},
		UpdateExpression: aws.String("SET #index = #index + :inc"),
		ExpressionAttributeNames: map[string]string{
			"#index": "index",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": types.AttributeValueMemberN{
				Value: "1",
			},
		},
		ReturnValues: types.ReturnValueUpdatedOld,
	})
	print(result)

	if err != nil {
		slog.ErrorCtx(ctx, "Error updating link index", err)
	}
	return 1, nil
}

func (r *DynamodbLinksRepository) SaveLink(ctx context.Context, link model.Link) error {
	return nil
}

func (r *DynamodbLinksRepository) GetLinkById(ctx context.Context, id string) (*model.Link, error) {
	return nil, nil
}

func (r *DynamodbLinksRepository) DeleteLinkById(ctx context.Context, id string) error {
	return nil
}

func (r *DynamodbLinksRepository) IncrementHitCount(ctx context.Context, id string) error {
	return nil
}
