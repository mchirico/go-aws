package sns

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/mchirico/go-aws/client"
)

func TestCreate(t *testing.T) {

	topic := "toprog3"
	input := &sns.CreateTopicInput{
		Name: &topic,
	}
	_, err := Create(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {

	topicArn := "toprog3"
	input := &sns.DeleteTopicInput{
		TopicArn: &topicArn,
	}
	_, err := Delete(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_FindARN(t *testing.T) {

	topicARN, err := FindARN(client.Config(), "toprog3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*topicARN)
}
