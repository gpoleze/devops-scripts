package ecr

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type EcrRepository struct {
	Name string
	URI  string
}

func getClient(region *string, profile *string) *ecr.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return ecr.NewFromConfig(cfg)
}

func DescribeRepositories(region *string, profile *string) []EcrRepository {
	client := getClient(region, profile)
	res, err := client.DescribeRepositories(context.TODO(), nil)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return nil
	}

	var repositories []EcrRepository

	for _, repository := range res.Repositories {
		ecrRepo := EcrRepository{
			Name: *repository.RepositoryName,
			URI:  *repository.RepositoryUri,
		}
		repositories = append(repositories, ecrRepo)
	}

	return repositories
}
