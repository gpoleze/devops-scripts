package iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListUsersInGroup(groupName, region, profile *string) ([]MyUser, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := iam.NewFromConfig(cfg)

	input := &iam.GetGroupInput{
		GroupName: groupName,
	}

	result, err := client.GetGroup(context.TODO(), input)
	if err != nil {
		fmt.Println("failed to get group users, %v", err)
		return nil, err
	}

	users := make([]MyUser, len(result.Users))
	for i, user := range result.Users {
		users[i] = NewMyUser(user)
	}
	return users, nil
}
