package main

import (
	"log"

	"github.com/gpoleze/devops-scripts/aws/ec2"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

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
	region, profile, outputType := utils.ReadAwsFlags()
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
