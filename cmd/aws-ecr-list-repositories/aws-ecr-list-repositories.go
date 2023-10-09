package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gpoleze/devops-scripts/aws/ecr"
	. "github.com/gpoleze/devops-scripts/table"
	"github.com/jedib0t/go-pretty/v6/table"
)

func readFlags() (*string, *string, *string) {
	var region string
	var profile string
	var outputType string
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&region, "r", "", "AWS region (shorthand)")

	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.StringVar(&profile, "p", "", "AWS profile (shorthand)")

	flag.StringVar(&outputType, "o", "table", "output type type(shorthand)")
	flag.StringVar(&outputType, "output", "table", "output type (accepted types table and json")

	flag.Parse()
	return &region, &profile, &outputType
}

func itemToRow(repository ecr.EcrRepository) table.Row {
	return table.Row{
		repository.Name,
		repository.URI,
	}
}

func printJson(repositories []ecr.EcrRepository) {
	val, _ := json.MarshalIndent(repositories, "", "    ")
	fmt.Println(string(val))
}

func main() {
	region, profile, outputType := readFlags()
	repositories := ecr.DescribeRepositories(region, profile)

	switch *outputType {
	case "json":
		printJson(repositories)
	case "table":
		BuildTable(repositories, itemToRow)
	default:
		BuildTable(repositories, itemToRow)

	}
}
