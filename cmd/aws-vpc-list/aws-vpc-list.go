package main

import (
	"github.com/gpoleze/devops-scripts/aws/ec2"
	. "github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func itemToTableRow(vpc ec2.Vpc) table.Row {
	return table.Row{
		vpc.Name,
		vpc.Id,
		vpc.CidrBlock,
	}
}

func main() {
	region, profile, outputType := ReadAwsFlags()
	vpcs := ec2.DescribeVpcs(region, profile)

	switch *outputType {
	case "json":
		PrintJson(vpcs)
	case "table":
		BuildTable(vpcs, itemToTableRow)
	default:
		BuildTable(vpcs, itemToTableRow)
	}

}
