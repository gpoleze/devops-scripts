package route53

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"log"
)

func ListRecords(zoneId string, region *string, profile *string) []types.ResourceRecordSet {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := route53.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListResourceRecordSets(context.TODO(), &route53.ListResourceRecordSetsInput{HostedZoneId: &zoneId})
	if err != nil {
		log.Fatal(err)
	}

	return output.ResourceRecordSets
}
