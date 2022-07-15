package sns

import (
	"fmt"

	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func Create(cfg aws.Config, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
	client := sns.NewFromConfig(cfg)
	return client.CreateTopic(context.TODO(), input)

}

func Delete(cfg aws.Config, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
	client := sns.NewFromConfig(cfg)
	return client.DeleteTopic(context.TODO(), input)

}

func FindARN(cfg aws.Config, topicName string) (*string, error) {
	client := sns.NewFromConfig(cfg)
	result, err := client.ListTopics(context.TODO(), &sns.ListTopicsInput{})
	if err != nil {
		return nil, err
	}
	for _, topic := range result.Topics {
		s := strings.Split(*topic.TopicArn, ":")
		if s[len(s)-1] == topicName {
			return topic.TopicArn, nil
		}

	}

	return nil, nil
}

func List(cfg aws.Config) error {
	client := sns.NewFromConfig(cfg)
	result, err := client.ListTopics(context.TODO(), &sns.ListTopicsInput{})
	if err != nil {
		return err
	}
	for _, topic := range result.Topics {
		fmt.Println(*topic.TopicArn)
		Attributes(cfg, *topic.TopicArn)
	}

	return nil
}

func Attributes(cfg aws.Config, topicArn string) error {
	client := sns.NewFromConfig(cfg)
	//client.SetTopicAttributes(ctx context.Context, params *sns.SetTopicAttributesInput, optFns ...func(*sns.Options))
	result, err := client.GetTopicAttributes(context.TODO(), &sns.GetTopicAttributesInput{TopicArn: &topicArn})

	if err != nil {
		return err
	}

	if val, ok := result.Attributes["Policy"]; ok {
		fmt.Println("\nPolicy:\n", val)
	}

	return nil
}

func Subs(cfg aws.Config, region, account, topic string) error {
	client := sns.NewFromConfig(cfg)
	arnTopic := fmt.Sprintf("arn:aws:sns:%s:%s:%s", region, account, topic)
	result, err := client.ListSubscriptionsByTopic(context.TODO(), &sns.ListSubscriptionsByTopicInput{TopicArn: &arnTopic})

	if err != nil {
		return err
	}
	subArn := ""
	for _, topic := range result.Subscriptions {
		fmt.Println(*topic.TopicArn)
		subArn = fmt.Sprintf("%s", *topic.SubscriptionArn)
	}

	if result, err := client.GetSubscriptionAttributes(context.TODO(),
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: &subArn}); err == nil {
		for k, v := range result.Attributes {
			fmt.Println(k, v)

		}
	}
	return nil
}

func Publish(cfg aws.Config, input *sns.PublishInput) (*sns.PublishOutput, error) {
	client := sns.NewFromConfig(cfg)

	return client.Publish(context.TODO(), input)

}

func Subscribe(cfg aws.Config, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	client := sns.NewFromConfig(cfg)

	return client.Subscribe(context.TODO(), input)

}
