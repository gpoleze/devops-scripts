package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gpoleze/devops-scripts/aws/rds"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
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

func main() {
	region, profile, outputType := readFlags()
	instances := rds.DescribeInstances(region, profile)

	var header table.Row

	for _, field := range reflect.VisibleFields(reflect.TypeOf(instances[0])) {
		header = append(header, field.Name)
	}

	if outputType == nil || *outputType == "table" {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(header)

		for _, instance := range instances {
			t.AppendRow(table.Row{
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
			})
		}

		t.Render()
		return
	}

	if *outputType == "json" {
		val, _ := json.MarshalIndent(instances, "", "    ")
		fmt.Println(string(val))
	}

}
