package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func CreateAccessKey(cfg aws.Config, userName string) (*iam.CreateAccessKeyOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.CreateAccessKeyInput{
		UserName: &userName,
	}
	return client.CreateAccessKey(context.TODO(), input)

}

func ListAccessKeys(cfg aws.Config, userName string) (*iam.ListAccessKeysOutput, error) {
	client := iam.NewFromConfig(cfg)
	var max int32 = 10
	input := &iam.ListAccessKeysInput{
		MaxItems: &max,
		UserName: &userName,
	}

	return client.ListAccessKeys(context.TODO(), input)

}

func DeleteAccessKey(cfg aws.Config, userName, keyId string) (*iam.DeleteAccessKeyOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.DeleteAccessKeyInput{
		AccessKeyId: &keyId,
		UserName:    &userName,
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

// General account
func GetAccountSummary(cfg aws.Config) (*iam.GetAccountSummaryOutput, error) {
	client := iam.NewFromConfig(cfg)
	input := &iam.GetAccountSummaryInput{}
	return client.GetAccountSummary(context.TODO(), input)

}

func ListAttachedGroupPolicies(cfg aws.Config, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.ListAttachedGroupPolicies(context.TODO(), input)
}

func ListAttachedRolePolicies(cfg aws.Config, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	client := iam.NewFromConfig(cfg)
	return client.ListAttachedRolePolicies(context.TODO(), input)
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
