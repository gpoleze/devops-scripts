package main

import (
	"fmt"
	"github.com/gpoleze/devops-scripts/aws/iam"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func main() {
	myFlags := append(utils.AwsFlags, utils.MyFlag{
		Name:        "group-name",
		Description: "IAM Group Name",
		ShortName:   "g",
		Required:    true,
	})

	parsedFlags, err := utils.ReadFlags(myFlags)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	users, err := iam.ListUsersInGroup(parsedFlags["group-name"], parsedFlags["region"], parsedFlags["profile"])

	if err != nil {
		fmt.Printf("Error found:\n%s\n", err)
		os.Exit(1)

	}
	switch *parsedFlags["output-type"] {
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
