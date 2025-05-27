package main

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/gpoleze/devops-scripts/aws/route53"
	. "github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func itemToTableRow(zone types.HostedZone) table.Row {
	return table.Row{
		*zone.Name,
		*zone.Id,
	}
}

func main() {
	region, profile, outputType := ReadAwsFlags()
	if *region == "" {
		*region = "us-east-1"
	}
	zones := route53.ListHostedZones(region, profile)

	switch *outputType {
	case "json":
		PrintJson(zones)
	case "table":
		BuildTableWithHeader(BuildTableParams[types.HostedZone]{zones, itemToTableRow, []string{"Name", "Id"}})
	default:
		BuildTable(zones, itemToTableRow)

	}

}
