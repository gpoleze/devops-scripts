package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gpoleze/devops-scripts/utils"
	"sort"
)

func DescribeVpcs(region, profile *string) Vpcs {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic(err)
	}
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVpcsInput{}
	result, err := client.DescribeVpcs(context.TODO(), input)
	if err != nil {
		panic(err)
	}

	var vpcs = Vpcs{}
	cidrs := []string{}
	for _, i := range result.Vpcs {
		vpc := Vpc{
			Name:      utils.FilterTagByKey(i.Tags, "Name"),
			Id:        *i.VpcId,
			CidrBlock: *i.CidrBlock,
		}
		vpcs = append(vpcs, vpc)
		cidrs = append(cidrs, *i.CidrBlock)
	}

	sort.Sort(vpcs)

	return vpcs

}
