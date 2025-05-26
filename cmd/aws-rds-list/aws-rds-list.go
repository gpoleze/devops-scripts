package main

import (
	"fmt"
	"github.com/gpoleze/devops-scripts/aws/rds"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func itemToTableRow(instance rds.RDSInstance) table.Row {
	return table.Row{
		instance.DBInstanceIdentifier,
		instance.DBInstanceStatus,
		instance.VpcId,
		instance.MasterUsername,
		instance.DBName,
		instance.DBInstanceClass,
		instance.Engine,
		instance.EngineVersion,
		instance.LatestRestorableTime,
		instance.InstanceCreateTime,
		instance.EndpointAddress,
	}
}

func main() {
	region, profile, outputType := utils.ReadAwsFlags()
	instances, err := rds.DescribeInstances(region, profile)

	if err != nil {
		fmt.Printf("Error found:\n%s\n", err)
		os.Exit(1)
	}

	switch *outputType {
	case "json":
		utils.PrintJson(&instances)
	case "table":
		utils.BuildTable(instances, itemToTableRow)
	default:
		utils.BuildTable(instances, itemToTableRow)

	}

}
