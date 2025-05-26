package utils

import "github.com/aws/aws-sdk-go-v2/service/ec2/types"

func FilterTagByKey(tags []types.Tag, keyName string) string {
	value := ""
	for _, tag := range tags {
		if *tag.Key == keyName {
			value = *tag.Value
		}
	}
	return value
}
