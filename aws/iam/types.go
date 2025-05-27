package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"time"
)

type MyUser struct {
	UserName         string
	UserId           string
	Arn              string
	CreateDate       string
	PasswordLastUsed string
}

func NewMyUser(user types.User) MyUser {
	myUser := MyUser{
		Arn:      *user.Arn,
		UserId:   *user.UserId,
		UserName: *user.UserName,
	}
	//if user.Arn == nil || *user.Arn == "" {
	//	print()
	//}
	//if user.UserId == nil || *user.UserId == "" {
	//	print()
	//}
	//if user.UserName == nil || *user.UserName == "" {
	//	print()
	//}
	if user.CreateDate == nil {
		myUser.CreateDate = ""
	} else {
		myUser.CreateDate = user.CreateDate.Format(time.RFC3339)
	}
	if user.PasswordLastUsed == nil {
		myUser.PasswordLastUsed = ""
	} else {
		myUser.PasswordLastUsed = user.PasswordLastUsed.Format(time.RFC3339)
	}
	return myUser
}
