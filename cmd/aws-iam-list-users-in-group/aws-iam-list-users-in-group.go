package main

import (
	"fmt"
	"github.com/gpoleze/devops-scripts/aws/iam"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type parameters struct {
	groupName  *string
	region     *string
	profile    *string
	outputType *string
}

func readParameters() parameters {
	var myFlags utils.MyFlags = append(utils.AwsFlags, utils.MyFlag{
		Name:        "group-name",
		Description: "IAM Group Name",
		ShortName:   "g",
		Required:    true,
	})

	myFlags.UpdateDefaultValue("region", "us-east-1")

	parsedFlags, err := utils.ReadFlags(myFlags)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return parameters{
		groupName:  parsedFlags["group-name"],
		region:     parsedFlags["region"],
		profile:    parsedFlags["profile"],
		outputType: parsedFlags["output-type"],
	}
}

func main() {
	parameters := readParameters()

	users, err := iam.ListUsersInGroup(parameters.groupName, parameters.region, parameters.profile)

	if err != nil {
		fmt.Printf("Error found:\n%s\n", err)
		os.Exit(1)

	}
	switch *parameters.outputType {
	case "json":
		utils.PrintJson(users)
	case "table":
		utils.BuildTable(users, itemToTableRow)
	default:
		utils.BuildTable(users, itemToTableRow)
	}
}

func itemToTableRow(user iam.MyUser) table.Row {
	return table.Row{
		user.UserName,
		user.UserId,
		user.Arn,
		user.CreateDate,
		user.PasswordLastUsed,
	}
}
