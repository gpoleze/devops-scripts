package main

import (
	"flag"
	"github.com/manifoldco/promptui"
	"log"
	"sort"

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

var repositoriesSelectTemplate = &promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "* {{ .Name | cyan }}",
	Inactive: "  {{ .Name | cyan }}",
	Selected: "* {{ .Name | red }} - {{ .URI | red }}",
	Details: `
{{ "Name:" | faint }}       {{ .Name }}
{{ "Id:" | faint }}         {{ .URI }}`,
}

func readFlags() arguments {
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

	//if repositoryName == "" {
	//	log.Fatal("The repository name cannot be empty")
	//}

	return arguments{
		Region:         &region,
		Profile:        &profile,
		OutputType:     &outputType,
		RepositoryName: &repositoryName,
	}
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

	if *arguments.RepositoryName == "" {
		repositories := ecr.DescribeRepositories(
			arguments.Region,
			arguments.Profile,
		)
		repository := selectItemFrom(&repositories, "Repositories", repositoriesSelectTemplate)
		*arguments.RepositoryName = repository.Name
	}

	images := ecr.DescribeImages(
		arguments.Region,
		arguments.Profile,
		arguments.RepositoryName,
	)

	if images == nil {
		log.Printf("The repository %s is empty", arguments.RepositoryName)
		return
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].PushedAt.Compare(images[j].PushedAt) > 0
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
