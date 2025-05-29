package main

import (
	"github.com/gpoleze/devops-scripts/aws/ec2"
	"os"
	"slices"
)

var (
	services = []string{"ec2", "s3", "ecr", "iam", "route53"}
	commands = map[string]map[string]func(){
		"ec2": {
			"list-vpcs":     ec2.ListVpcs,
			"describe-vpcs": ec2.ListVpcs,
		},
	}
)

func main() {
	args := os.Args
	if len(args) < 2 {
		os.Exit(1)
	}

	if !slices.Contains(services, args[1]) {
		os.Exit(1)
	}
	println(args)
	service := args[1]
	command := args[2]

	commands[service][command]()
}
