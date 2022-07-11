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

func DeleteAccessKey(cfg aws.Config, userName, keyId string) (*iam.DeleteAccessKeyOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.DeleteAccessKeyInput{
		AccessKeyId: new(string),
		UserName:    new(string),
	}
	return client.DeleteAccessKey(context.TODO(), input)

}

func AddUserToGroup(cfg aws.Config, userName, groupName string) (*iam.AddUserToGroupOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.AddUserToGroupInput{
		GroupName: &groupName,
		UserName:  &userName,
	}

	return client.AddUserToGroup(context.TODO(), input)

}

func ListUsers(cfg aws.Config, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.ListUsers(context.TODO(), input)

}

func CreateUser(cfg aws.Config, input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.CreateUser(context.TODO(), input)
}

func DeleteUser(cfg aws.Config, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.DeleteUser(context.TODO(), input)
}

func RemoveUserFromGroup(cfg aws.Config, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.RemoveUserFromGroup(context.TODO(), input)
}
