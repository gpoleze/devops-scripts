package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/gpoleze/devops-scripts/aws/ecr"
	"github.com/gpoleze/devops-scripts/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
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
	arguments := readArguments()

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

func readArguments() arguments {
	flags := append(utils.AwsFlags, utils.MyFlag{
		Name:        "repository",
		ShortName:   "n",
		Description: "The name of the ECR repository to list images from.",
	})
	parsedFlags, err := utils.ReadFlags(flags)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	arguments := arguments{
		Region:         parsedFlags["region"],
		Profile:        parsedFlags["profile"],
		RepositoryName: parsedFlags["repository"],
		OutputType:     parsedFlags["output-type"],
	}
	return arguments
}
