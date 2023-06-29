package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang.org/x/exp/slog"

	"gotiny/internal/core/model"
	"gotiny/internal/data/dto"
	"gotiny/internal/data/util"
)

const (
	linksGlobalDataPK  = "LINKS_GLOBAL"
	linksGlobalIndexSK = "INDEX"
	indexKey           = "Index"
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
) (*DynamodbLinksRepository, error) {
	r := DynamodbLinksRepository{client: client, config: cfg}
	err := r.init()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *DynamodbLinksRepository) GetNextLinkIndex(ctx context.Context) (uint, error) {
	result, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(r.config.LinksTableName()),
		Key:              util.FormatKeys(linksGlobalDataPK, linksGlobalIndexSK),
		UpdateExpression: aws.String("SET #index = #index + :inc"),
		ExpressionAttributeNames: map[string]string{
			"#index": indexKey,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": &types.AttributeValueMemberN{
				Value: "1",
			},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error updating link index", "err", err)
		return 0, err
	}

	var data dto.GlobalLinksDataDto
	err = attributevalue.UnmarshalMap(result.Attributes, &data)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling link index", "err", err)
		return 0, err
	}

	return data.Index, nil
}

func (r *DynamodbLinksRepository) SaveLink(ctx context.Context, link model.Link) error {
	dto := dto.LinkDtoFromLink(link)
	avs, err := attributevalue.MarshalMap(dto)
	if err != nil {
		slog.ErrorCtx(ctx, "Error marshalling link", "err", err)
		return err
	}
	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Item:      avs,
	})

	if err != nil {
		slog.ErrorCtx(ctx, "Error saving link", "err", err)
		return err
	}

	return nil
}

func (r *DynamodbLinksRepository) GetLinkById(ctx context.Context, id string) (*model.Link, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Key:       util.FormatKeys(dto.LinkPK, dto.LinkSKPrefix+id),
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error getting link by id", "err", err)
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	var dto dto.LinkDto
	err = attributevalue.UnmarshalMap(result.Item, &dto)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling link", "err", err)
		return nil, err
	}
	link := dto.ToLink()

	return &link, nil
}

func (r *DynamodbLinksRepository) DeleteLinkById(ctx context.Context, id string) error {
	_, err := r.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Key:       util.FormatKeys(dto.LinkPK, dto.LinkSKPrefix+id),
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error deleting link by id", "err", err)
		return err
	}

	return nil
}

func (r *DynamodbLinksRepository) IncrementHitCount(ctx context.Context, id string) error {
	_, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(r.config.LinksTableName()),
		Key:              util.FormatKeys(dto.LinkPK, dto.LinkSKPrefix+id),
		UpdateExpression: aws.String("SET #hits = #hits + :inc"),
		ExpressionAttributeNames: map[string]string{
			"#hits": "Hits",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": &types.AttributeValueMemberN{
				Value: "1",
			},
		},
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error incrementing hit count", "err", err)
		return err
	}

	return nil
}

func (r *DynamodbLinksRepository) init() error {
	result, err := r.client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Key:       util.FormatKeys(linksGlobalDataPK, linksGlobalIndexSK),
	})
	if err != nil {
		slog.ErrorCtx(context.Background(), "Initialization error - get item", "err", err)
		return err
	}
	if result.Item != nil {
		return nil
	}

	initialData := dto.GlobalLinksDataDto{
		PK:    linksGlobalDataPK,
		SK:    linksGlobalIndexSK,
		Index: 0,
	}
	formatedData, err := attributevalue.MarshalMap(initialData)
	if err != nil {
		slog.ErrorCtx(context.Background(), "Initialization error - marshal map", "err", err)
		return err
	}

	_, err = r.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Item:      formatedData,
	})

	return nil
}
