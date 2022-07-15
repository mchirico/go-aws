package sns

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/mchirico/go-aws/client"
)

func TestCreate(t *testing.T) {

	topic := "sns-to-lambda"
	input := &sns.CreateTopicInput{
		Name: &topic,
	}
	_, err := Create(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublish(t *testing.T) {

	topic := "sns-to-lambda"
	topicARN, err := FindARN(client.Config(), topic)
	subject := "test-sns-to-lambda"
	message := "test-message-10"
	input := &sns.PublishInput{
		Message:  &message,
		Subject:  &subject,
		TopicArn: topicARN,
	}
	_, err = Publish(client.Config(), input)
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

func TestSubscribe(t *testing.T) {

	topic := "sns-to-lambda"
	topicARN, err := FindARN(client.Config(), topic)
	endPoint := "arn:aws:lambda:us-east-1:652115786209:function:sns"
	protocol := "LAMBDA"
	input := &sns.SubscribeInput{
		Protocol:              &protocol,
		TopicArn:              topicARN,
		Attributes:            map[string]string{},
		Endpoint:              &endPoint,
		ReturnSubscriptionArn: true,
	}
	result, err := Subscribe(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_FindARN(t *testing.T) {

	topicARN, err := FindARN(client.Config(), "toprog3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*topicARN)
}
