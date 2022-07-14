/*
iam_cmb ... iam combination. FIXME (mmc), come up with
a better name.

*/

package iam

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/mchirico/go-aws/client"
)

type I struct {
	userName string
	cfg      aws.Config
	max      int32
}

func NewI(userName string) *I {
	cfg := client.Config()
	return &I{userName: userName, max: 50, cfg: cfg}
}

func (i *I) CreateAccessKey() (*iam.CreateAccessKeyOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	input := &iam.CreateAccessKeyInput{
		UserName: &i.userName,
	}
	return client.CreateAccessKey(context.TODO(), input)

}

func (i *I) ListAccessKeys() (*iam.ListAccessKeysOutput, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.ListAccessKeysInput{
		MaxItems: &i.max,
		UserName: &i.userName,
	}

	return client.ListAccessKeys(context.TODO(), input)

}

func (i *I) ListAttachedGroupPolicies(group string) (*iam.ListAttachedGroupPoliciesOutput, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.ListAttachedGroupPoliciesInput{
		GroupName: &group,
		MaxItems:  &i.max,
	}

	return client.ListAttachedGroupPolicies(context.TODO(), input)

}

func (i *I) ListRoles() (*iam.ListRolesOutput, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.ListRolesInput{
		MaxItems: &i.max,
	}
	return client.ListRoles(context.TODO(), input)
}

func (i *I) GetRole(role string) (*string, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.ListRolesInput{
		MaxItems: &i.max,
	}
	result, err := client.ListRoles(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	for _, v := range result.Roles {
		if *v.RoleName == role {
			return v.Arn, nil
		}
	}
	return nil, errors.New("not found")
}

func (i *I) AttachRolePolicy(roleName, policyArn string) (*iam.AttachRolePolicyOutput, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.AttachRolePolicyInput{
		PolicyArn: &policyArn,
		RoleName:  &roleName,
	}
	return client.AttachRolePolicy(context.TODO(), input)

}

func (i *I) DetachRolePolicy(roleName, policyArn string) (*iam.DetachRolePolicyOutput, error) {
	client := iam.NewFromConfig(i.cfg)

	input := &iam.DetachRolePolicyInput{
		PolicyArn: &policyArn,
		RoleName:  &roleName,
	}
	return client.DetachRolePolicy(context.TODO(), input)

}

func (i *I) CreateRole(input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	return client.CreateRole(context.TODO(), input)

}

func (i *I) createUser(input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	return client.CreateUser(context.TODO(), input)
}

func (i *I) AccessReport() (*iam.GetCredentialReportOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	input := &iam.GetCredentialReportInput{}

	return client.GetCredentialReport(context.TODO(), input)

}

func WriteFile(file string, data []byte) error {
	return os.WriteFile(file, data, 0644)
}
