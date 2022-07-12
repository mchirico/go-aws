/*
iam_cmb ... iam combination. FIXME (mmc), come up with
a better name.

*/

package iam

import (
	"context"
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

func (i *I) createUser(input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	return client.CreateUser(context.TODO(), input)
}

func WriteFile(file string, data []byte) error {
	return os.WriteFile(file, data, 0644)
}
