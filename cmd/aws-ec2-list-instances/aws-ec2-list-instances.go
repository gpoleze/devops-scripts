package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gpoleze/devops-scripts/aws/ec2"
	"github.com/gpoleze/devops-scripts/utils"
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

	fmt.Printf("===============> regiao: %s", region)

	if region == "" {
		log.Fatal("O parametro --region é obrigatório")
	}

	return &region, &profile, &outputType
}

func itemToTableRow() func(instance ec2.MyInstanceInfo) table.Row {
	return func(instance ec2.MyInstanceInfo) table.Row {
		return table.Row{
			instance.Name,
			instance.Id,
			instance.Type,
			instance.State,
			instance.Ami,
			instance.LaunchTime,
			instance.PrivateIp,
			instance.PublicIp,
		}
	}
}

func main() {
	region, profile, outputType := readFlags()
	instances := ec2.DescribeInstances(region, profile)
	if instances == nil {
		log.Fatal("No instances found")
	}

	switch *outputType {
	case "json":
		utils.PrintJson(instances)
	case "table":
		utils.BuildTable(instances, itemToTableRow())
	default:
		utils.BuildTable(instances, itemToTableRow())
	}

}
