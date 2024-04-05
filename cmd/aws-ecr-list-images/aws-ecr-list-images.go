package main

import (
	"flag"
	"fmt"
	"sort"
	"errors"
	
	"github.com/gpoleze/devops-scripts/aws/ecr"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

type arguments struct {
	Region         *string
	Profile        *string
	OutputType     *string
	RepositoryName *string
}

func readFlags() (arguments, error) {
	var region string
	var profile string
	var outputType string
	var repositoryName string

	flag.StringVar(&region, "region", "", "AWS region")
	flag.StringVar(&region, "r", "", "AWS region (shorthand)")

	flag.StringVar(&profile, "profile", "", "AWS profile")
	flag.StringVar(&profile, "p", "", "AWS profile (shorthand)")

	flag.StringVar(&outputType, "output", "table", "output type (accepted types table and json")
	flag.StringVar(&outputType, "o", "table", "output type type(shorthand)")

	flag.StringVar(&repositoryName, "repository-name", "", "ECR Repository Name")
	flag.StringVar(&repositoryName, "n", "", "ECR Repository Name (shorthand)")

	flag.Parse()

	if repositoryName == "" {
		return arguments{}, errors.New("the repository name cannot be empty")
	}
	if region == "" {
		return arguments{}, errors.New("the region cannot be empty")
	}

	return arguments{
		Region:         &region,
		Profile:        &profile,
		OutputType:     &outputType,
		RepositoryName: &repositoryName,
	},nil
}

func itemToTableRow(image ecr.EcrImage) table.Row {

	return table.Row{
		image.Tags,
		image.PushedAt,
		image.MediaType,
		image.SizeMB,
		image.Digest,
	}

}

func main() {
	arguments,err := readFlags()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	images := ecr.DescribeImages(
		arguments.Region,
		arguments.Profile,
		arguments.RepositoryName,
	)

	if len(images) == 0 {
		fmt.Printf("No images found on repository '%s' in '%s' \n", *arguments.RepositoryName, *arguments.Region)
		return
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].PushedAt.After(images[j].PushedAt)
	})

	switch *arguments.OutputType {
	case "json":
		utils.PrintJson(images)
	case "table":
		utils.BuildTable(images, itemToTableRow)
	default:
		utils.BuildTable(images, itemToTableRow)
	}
}
