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

type PrimaryKeys struct {
	PK string
	SK string
}

type Gsi1Keys struct {
	GSI_1_PK string
	GSI_1_SK string
}

func EncodePageToken(at map[string]types.AttributeValue) (*string, error) {
	if at == nil {
		return nil, nil
	}
	var keys map[string]interface{}
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

func DecodePageToken(token *string) (map[string]types.AttributeValue, error) {
	if token == nil {
		return nil, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(*token)
	if err != nil {
		return nil, err
	}
	var keys map[string]interface{}
	err = json.Unmarshal(decoded, &keys)
	if err != nil {
		return nil, err
	}
	ats, err := attributevalue.MarshalMap(keys)

	return ats, err
}
