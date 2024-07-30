package ecr

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"strings"
	"time"
)

type EcrImage struct {
	Tags      string
	PushedAt  time.Time
	MediaType string
	SizeMB    int64
	Digest    string
}

func DescribeImages(region *string, profile *string, repositoryName *string) []EcrImage {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ecr.NewFromConfig(cfg)

	params := ecr.DescribeImagesInput{RepositoryName: repositoryName}

	res, err := client.DescribeImages(context.TODO(), &params)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return nil
	}

	var images []EcrImage

	if len(res.ImageDetails) == 0
	  return nil

	for _, image := range res.ImageDetails {
		ecrImage := EcrImage{
			Tags:     strings.Join(image.ImageTags, ", "),
			PushedAt: *image.ImagePushedAt,
			SizeMB:   *image.ImageSizeInBytes / 1024 / 1024,
			Digest:   *image.ImageDigest,
		}
		if image.ArtifactMediaType != nil && *image.ArtifactMediaType == "application/vnd.docker.container.image.v1+json" {
			ecrImage.MediaType = "image"
		} else {
			ecrImage.MediaType = "Image Index"
		}

		images = append(images, ecrImage)
	}

	return images
}
