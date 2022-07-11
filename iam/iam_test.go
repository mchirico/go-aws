package iam

import (
	"errors"
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

	result, err := CreateAccessKey(client.Config(), userName)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func TestListKey(t *testing.T) {
	userName := "go-aws"

	result, err := ListAccessKeys(client.Config(), userName)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.AccessKeyMetadata {
		fmt.Println(*v.AccessKeyId)
	}

}

func GetKeyId(userName string) (string, error) {
	result, err := ListAccessKeys(client.Config(), userName)
	if err != nil {
		return "", err
	}
	for _, v := range result.AccessKeyMetadata {
		return *v.AccessKeyId, nil
	}
	return "", errors.New("No key found")
}

func TestDeleteKey(t *testing.T) {
	userName := "go-aws"
	key, err := GetKeyId(userName)
	if err != nil {
		t.Fatal(err)
	}
	result, err := DeleteAccessKey(client.Config(), userName, key)
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

func TestGetAccountSummary(t *testing.T) {

	result, err := GetAccountSummary(client.Config())
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range result.SummaryMap {
		fmt.Println(k, v)
	}

}
