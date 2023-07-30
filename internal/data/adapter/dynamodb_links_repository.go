package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/segmentio/ksuid"
	"golang.org/x/exp/slog"

	"gotiny/internal/config"
	"gotiny/internal/core/model"
	core_util "gotiny/internal/core/util"
	"gotiny/internal/data/dto"
	"gotiny/internal/data/util"
)

const (
	linksGlobalDataPK  = "LINKS_GLOBAL"
	linksGlobalIndexSK = "INDEX"
	indexKey           = "Index"
	pageLimit          = 30
)

type DynamodbLinksRepository struct {
	client *dynamodb.Client
	config *config.Config
}

func NewDynamodbLinksRepository(
	client *dynamodb.Client,
	cfg *config.Config,
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
	var d dto.LinkDto
	err = attributevalue.UnmarshalMap(result.Item, &d)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling link", "err", err)
		return nil, err
	}
	link := dto.LinkDtoToLink(d)

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

func (r *DynamodbLinksRepository) SaveHitAnalitics(
	ctx context.Context,
	linkId string,
	requestData model.LinkHitAnalitics,
) error {
	ttl := uint(time.Now().Add(time.Hour * 24 * 30).Unix()) // Valid for 30 days
	dto := dto.LinkHitAnaliticsDto{
		PK:          dto.LinkVisitPKPrefix + linkId,
		SK:          dto.LinkVisitSKPrefix + ksuid.New().String(),
		IpDetails:   requestData.IpDetails,
		RequestData: requestData.RequestData,
		CreatedAt:   time.Now(),
		TTL:         &ttl,
	}
	avs, err := attributevalue.MarshalMap(dto)
	if err != nil {
		slog.ErrorCtx(ctx, "Error marshalling link visit", "err", err)
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.config.LinksTableName()),
		Item:      avs,
	})

	if err != nil {
		slog.ErrorCtx(ctx, "Error saving link visit", "err", err)
		return err
	}

	return nil
}

func (r *DynamodbLinksRepository) GetLinkVisits(
	ctx context.Context,
	linkId string,
	page *string,
) (model.PagedResponse[model.LinkHitAnalitics], error) {
	lastEvaluatedKey, err := util.DecodePageToken(page)
	if err != nil {
		slog.ErrorCtx(ctx, "Error decoding page token", "err", err)
		return model.PagedResponse[model.LinkHitAnalitics]{}, err
	}
	result, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.config.LinksTableName()),
		KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: dto.LinkVisitPKPrefix + linkId},
			":sk": &types.AttributeValueMemberS{Value: dto.LinkVisitSKPrefix},
		},
		Limit:             aws.Int32(pageLimit),
		ExclusiveStartKey: lastEvaluatedKey,
		ScanIndexForward:  aws.Bool(false),
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error getting link visits", "err", err)
		return model.PagedResponse[model.LinkHitAnalitics]{}, err
	}

	var visits []dto.LinkHitAnaliticsDto
	err = attributevalue.UnmarshalListOfMaps(result.Items, &visits)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling link visits", "err", err)
		return model.PagedResponse[model.LinkHitAnalitics]{}, err
	}
	pageToken, err := util.EncodePageToken(result.LastEvaluatedKey)
	if err != nil {
		slog.ErrorCtx(ctx, "Error decoding page token", "err", err)
		return model.PagedResponse[model.LinkHitAnalitics]{}, err
	}

	return model.PagedResponse[model.LinkHitAnalitics]{
		Items: core_util.MapSlice[dto.LinkHitAnaliticsDto, model.LinkHitAnalitics](
			visits,
			dto.LinkHitAnaliticsDtoToDomain,
		),
		PageToken: pageToken,
	}, nil
}

func (r *DynamodbLinksRepository) GetUserLinks(
	ctx context.Context,
	userId string,
	page *string,
) (model.PagedResponse[model.Link], error) {
	lastEvaluatedKey, err := util.DecodePageToken(page)
	if err != nil {
		slog.ErrorCtx(ctx, "Error decoding page token", "err", err)
		return model.PagedResponse[model.Link]{}, err
	}
	result, err := r.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(r.config.LinksTableName()),
		IndexName:              aws.String("GSI_1"),
		KeyConditionExpression: aws.String("GSI_1_PK = :pk AND begins_with(GSI_1_SK, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: dto.LinkGSI1PKPrefix + userId},
			":sk": &types.AttributeValueMemberS{Value: dto.LinkGSI1SKPrefix},
		},
		Limit:             aws.Int32(pageLimit),
		ExclusiveStartKey: lastEvaluatedKey,
		ScanIndexForward:  aws.Bool(false),
	})
	if err != nil {
		slog.ErrorCtx(ctx, "Error getting user links", "err", err)
		return model.PagedResponse[model.Link]{}, err
	}
	var links []dto.LinkDto
	err = attributevalue.UnmarshalListOfMaps(result.Items, &links)
	if err != nil {
		slog.ErrorCtx(ctx, "Error unmarshalling user links", "err", err)
		return model.PagedResponse[model.Link]{}, err
	}
	fmt.Printf("LastEvaluatedKey: %v\n", result.LastEvaluatedKey)
	pageToken, err := util.EncodePageToken(result.LastEvaluatedKey)
	if err != nil {
		slog.ErrorCtx(ctx, "Error decoding page token", "err", err)
		return model.PagedResponse[model.Link]{}, err
	}

	return model.PagedResponse[model.Link]{
		Items: core_util.MapSlice[dto.LinkDto, model.Link](
			links,
			dto.LinkDtoToLink,
		),
		PageToken: pageToken,
	}, nil
}
