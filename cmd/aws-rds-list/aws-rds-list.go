package main

import (
	"github.com/gpoleze/devops-scripts/aws/rds"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
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
	instances := rds.DescribeInstances(region, profile)

	switch *outputType {
	case "json":
		utils.PrintJson(&instances)
	case "table":
		utils.BuildTable(instances, itemToTableRow)
	default:
		utils.BuildTable(instances, itemToTableRow)

	}

}
