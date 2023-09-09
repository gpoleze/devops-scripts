package main

import (
	"flag"
	"github.com/gpoleze/devops-scripts/pkg/aws/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
)

func _map(list *[]any, funk func(any) any) []any {
	var new_list []any
	for _, field := range *list {
		new_list = append(new_list, funk(field))
	}
	return new_list
}

func main() {

	var region string
	var profile string
	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&region, "r", "", "AWS region (shorthand)")

	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.StringVar(&profile, "p", "", "AWS profile (shorthand)")

	flag.Parse()

	instances := ec2.DescribeInstances(&region, &profile)

	var header table.Row

	for _, field := range reflect.VisibleFields(reflect.TypeOf(instances[0])) {
		header = append(header, field.Name)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(header)

	for _, instance := range instances {
		t.AppendRow(table.Row{
			instance.Name,
			instance.Id,
			instance.Type,
			instance.State,
			instance.Ami,
			instance.LaunchTime,
			instance.PrivateIp,
			instance.PublicIp,
		})
	}

	t.Render()
}
