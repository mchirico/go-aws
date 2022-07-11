package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func CreateKey(cfg aws.Config, userName string) (*iam.CreateAccessKeyOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.CreateAccessKeyInput{
		UserName: &userName,
	}
	return client.CreateAccessKey(context.TODO(), input)

}

func ListUsers(cfg aws.Config, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.ListUsers(context.TODO(), input)

}

func CreateUser(cfg aws.Config, input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.CreateUser(context.TODO(), input)
}
