package ec2

import (
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func itemToTableRow(vpc Vpc) table.Row {
	return table.Row{
		vpc.Name,
		vpc.Id,
		vpc.CidrBlock,
	}
}

func ListVpcs() {
	region, profile, outputType := utils.ReadAwsFlags()
	vpcs := DescribeVpcs(region, profile)

	switch *outputType {
	case "json":
		utils.PrintJson(vpcs)
	case "table":
		utils.BuildTable(vpcs, itemToTableRow)
	default:
		utils.BuildTable(vpcs, itemToTableRow)
	}
}
