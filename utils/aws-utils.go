package utils

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/jedib0t/go-pretty/v6/table"
)

func FilterTagByKey(tags []types.Tag, keyName string) string {
	value := ""
	for _, tag := range tags {
		if *tag.Key == keyName {
			value = *tag.Value
		}
	}
	return value
}

func PrintOutput[T any](output string, list []T, itemToTableRow func(T) table.Row) {
	switch output {
	case "json":
		PrintJson(list)
	case "table":
		BuildTable(list, itemToTableRow)
	default:
		BuildTable(list, itemToTableRow)
	}
}
