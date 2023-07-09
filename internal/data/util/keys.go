package util

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func FormatKeys(PK string, SK string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{
			Value: PK,
		},
		"SK": &types.AttributeValueMemberS{
			Value: SK,
		},
	}
}

type keys struct {
	PK string
	SK string
}

func EncodePrimaryPageToken(at map[string]types.AttributeValue) (*string, error) {
	if at == nil {
		return nil, nil
	}
	var keys keys
	err := attributevalue.UnmarshalMap(at, &keys)
	if err != nil {
		return nil, err
	}
	jsonEncoded, err := json.Marshal(keys)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(jsonEncoded)

	return &encoded, nil
}

func DecodePrimaryPageToken(token *string) (map[string]types.AttributeValue, error) {
	if token == nil {
		return nil, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(*token)
	if err != nil {
		return nil, err
	}
	var keys keys
	err = json.Unmarshal(decoded, &keys)
	if err != nil {
		return nil, err
	}
	ats, err := attributevalue.MarshalMap(keys)

	return ats, err
}
