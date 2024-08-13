package main

import (
	"flag"
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

func readFlags() (arguments, error) {
	var region string
	var profile string
	var outputType string
	var hostedZoneName string

	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&region, "r", "", "AWS region (shorthand)")

	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.StringVar(&profile, "p", "", "AWS profile (shorthand)")

	flag.StringVar(&outputType, "output", "table", "output type (accepted types table and json")
	flag.StringVar(&outputType, "o", "table", "output type type(shorthand)")

	flag.StringVar(&hostedZoneName, "hosted-zone-name", "", "Hosted Zone Name")
	flag.StringVar(&hostedZoneName, "n", "", "Hosted Zone Name (shorthand)")

	flag.Parse()

	return arguments{
		Region:         &region,
		Profile:        &profile,
		OutputType:     &outputType,
		HostedZoneName: &hostedZoneName,
	}, nil
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
	arguments, err := readFlags()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if *arguments.Region == "" {
		*arguments.Region = "us-east-1"
	}

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
