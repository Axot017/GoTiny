package util

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

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
