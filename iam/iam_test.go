package iam

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
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

func TestCreateUsers(t *testing.T) {
	userName := "go-aws"
	key := "Name"
	value := "go-aws"
	tag := types.Tag{Key: &key, Value: &value}
	input := &iam.CreateUserInput{
		UserName: &userName,
		Tags:     []types.Tag{tag},
	}
	result, err := CreateUser(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func TestCreateKey(t *testing.T) {
	userName := "go-aws"

	result, err := CreateKey(client.Config(), userName)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func TestAddUserToGroup(t *testing.T) {
	userName := "go-aws"
	groupName := "Admin"

	result, err := AddUserToGroup(client.Config(), userName, groupName)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func Test_RemoveUserFromGroup(t *testing.T) {

	userName := "go-aws"
	groupName := "Admin"
	input := &iam.RemoveUserFromGroupInput{
		GroupName: &groupName,
		UserName:  &userName,
	}
	result, err := RemoveUserFromGroup(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func TestDeleteUser(t *testing.T) {
	userName := "go-aws"

	input := &iam.DeleteUserInput{
		UserName: &userName,
	}
	result, err := DeleteUser(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}
