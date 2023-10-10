package main

import (
	"github.com/gpoleze/devops-scripts/aws/ecr"
	. "github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func itemToTableRow(repository ecr.EcrRepository) table.Row {
	return table.Row{
		repository.Name,
		repository.URI,
	}
}

func main() {
	region, profile, outputType := ReadAwsFlags()
	repositories := ecr.DescribeRepositories(region, profile)

	switch *outputType {
	case "json":
		PrintJson(repositories)
	case "table":
		BuildTable(repositories, itemToTableRow)
	default:
		BuildTable(repositories, itemToTableRow)

	}
}
