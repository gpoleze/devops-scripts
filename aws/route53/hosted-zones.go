package route53

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"log"
	"strings"
)

func ListHostedZones(region *string, profile *string) []types.HostedZone {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := route53.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListHostedZones(context.TODO(), &route53.ListHostedZonesInput{})
	if err != nil {
		log.Fatal(err)
	}

	return output.HostedZones
}

func GetHostedZoneByName(zoneName string, region *string, profile *string) (zone types.HostedZone, error error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := route53.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListHostedZones(context.TODO(), &route53.ListHostedZonesInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, zone := range output.HostedZones {
		if strings.Compare(*zone.Name, zoneName) == 0 {
			return zone, nil
		}
	}

	return zone, errors.New("no zone with name " + zoneName + " found")
}
