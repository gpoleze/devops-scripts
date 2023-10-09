package rds

import (
	"context"
	"fmt"
	"time"

	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDSInstance struct {
	DBInstanceIdentifier string
	DBInstanceStatus     string
	VpcId                string
	MasterUsername       string
	DBName               string
	DBInstanceClass      string
	Engine               string
	EngineVersion        string
	LatestRestorableTime time.Time
	InstanceCreateTime   time.Time
	EndpointAddress      string
}

func DescribeInstances(region *string, profile *string) []RDSInstance {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return nil
	}
	rdsClient := rds.NewFromConfig(cfg)

	output, err := rdsClient.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})

	if err != nil {
		fmt.Printf("Couldn't list DB instances: %v\n", err)
		return nil
	}

	var instances = []RDSInstance{}
	for _, instanceOutput := range output.DBInstances {
		var instance = RDSInstance{
			DBInstanceIdentifier: *instanceOutput.DBInstanceIdentifier,
			DBInstanceStatus:     *instanceOutput.DBInstanceStatus,
			VpcId:                *instanceOutput.DBSubnetGroup.VpcId,
			MasterUsername:       *instanceOutput.MasterUsername,
			DBInstanceClass:      *instanceOutput.DBInstanceClass,
			Engine:               *instanceOutput.Engine,
			EngineVersion:        *instanceOutput.EngineVersion,
			EndpointAddress:      *instanceOutput.Endpoint.Address,
		}

		if instanceOutput.DBName != nil {
			instance.DBName = *instanceOutput.DBName
		}
		if instanceOutput.InstanceCreateTime != nil {
			instance.InstanceCreateTime = *instanceOutput.InstanceCreateTime
		}
		if instanceOutput.LatestRestorableTime != nil {
			instance.LatestRestorableTime = *instanceOutput.LatestRestorableTime
		}

		instances = append(instances, instance)

	}

	return instances

}
