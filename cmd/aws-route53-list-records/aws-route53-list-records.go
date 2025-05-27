package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/gpoleze/devops-scripts/aws/route53"
	. "github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

type arguments struct {
	Region         *string
	Profile        *string
	OutputType     *string
	HostedZoneName *string
}

type myZone struct {
	Name string
	Id   string
}

func itemToTableRow(record types.ResourceRecordSet) table.Row {
	var resourceRecords []string
	for _, r := range record.ResourceRecords {
		resourceRecords = append(resourceRecords, *r.Value)
	}

	return table.Row{
		*record.Name,
		record.Type,
		strings.Join(resourceRecords, ","),
	}
}

func readFlags() arguments {
	flags := append(AwsFlags, MyFlag{
		Name:        "hosted-zone-name",
		ShortName:   "n",
		Description: "Hosted Zone Name",
	})
	parsedFlags, err := ReadFlags(flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	arguments := arguments{
		Region:         parsedFlags["region"],
		Profile:        parsedFlags["profile"],
		HostedZoneName: parsedFlags["hosted-zone-name"],
		OutputType:     parsedFlags["output-type"],
	}

	if *arguments.Region == "" {
		*arguments.Region = "us-east-1"
	}

	return arguments
}

var zoneSelectTemplate = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "* {{ .Name | cyan }}",
	Inactive: "  {{ .Name | cyan }}",
	Selected: "* {{ .Name | red }} - {{ .ID | red }}",
	Details: `
{{ "Name:" | faint }}       {{ .Name }}
{{ "Id:" | faint }}         {{ .Id }}`,
}

func selectItemFrom[T any](items *[]T, label string, template *promptui.SelectTemplates) *T {
	prompt := promptui.Select{
		Label:     label,
		Items:     *items,
		Templates: template,
		Size:      8,
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return &(*items)[index]
}

func main() {
	arguments := readFlags()
	var zoneId string

	if *arguments.HostedZoneName != "" {
		zone, err := route53.GetHostedZoneByName(*arguments.HostedZoneName, arguments.Region, arguments.Profile)
		if err != nil {
			log.Fatal(err)
		}
		zoneId = *zone.Id
	} else {
		zones := route53.ListHostedZones(arguments.Region, arguments.Profile)
		var myZones []myZone
		for _, z := range zones {
			myZones = append(myZones, myZone{*z.Name, *z.Id})
		}

		zoneId = selectItemFrom(&myZones, "Select Zones", zoneSelectTemplate).Id

	}
	records := route53.ListRecords(zoneId, arguments.Region, arguments.Profile)

	switch *arguments.OutputType {
	case "json":
		PrintJson(records)
	case "table":
		BuildTableWithHeader(BuildTableParams[types.ResourceRecordSet]{records, itemToTableRow, []string{"Name", "Type", "Records"}})
	default:
		BuildTableWithHeader(BuildTableParams[types.ResourceRecordSet]{records, itemToTableRow, []string{"Name", "Type", "Records"}})
	}
}
