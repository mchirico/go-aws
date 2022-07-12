/*
iam_cmb ... iam combination. FIXME (mmc), come up with
a better name.

*/

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type I struct {
	userName string
	cfg      aws.Config
	max      int32
}

func NewI(userName string) *I {
	return &I{userName: userName, max: 50}
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

func (i *I) createUser(input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	client := iam.NewFromConfig(i.cfg)
	return client.CreateUser(context.TODO(), input)
}
