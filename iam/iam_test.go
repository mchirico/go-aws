package iam

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/mchirico/go-aws/client"
)

func TestListUsers(t *testing.T) {
	var max int32 = 20
	input := &iam.ListUsersInput{
		MaxItems: &max,
	}
	result, err := ListUsers(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.Users {
		fmt.Println(*v.UserName)
	}

}
