package main

import (
	"flag"
	"github.com/thoas/go-funk"
	"gitlab.com/gabriel.poleze/ssh-aws/aws/ec2"
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
	header := funk.Map(reflect.VisibleFields(reflect.TypeOf(instances[0])), func(field reflect.StructField) string {
		return field.Name
	})

	println("%s", header)

	//t := table.NewWriter()
	//t.SetOutputMirror(os.Stdout)
	//t.AppendHeader(table.Row{header})
	//
	//t.Render()
}
