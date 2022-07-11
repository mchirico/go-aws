package vpc

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func DescribeTypeOfferings(cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstanceTypeOfferingsInput{}
	result, err := client.DescribeInstanceTypeOfferings(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil

}

func DescribeInstances(cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstanceTypeOfferingsInput{}
	result, err := client.DescribeInstanceTypeOfferings(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil

}
func DescribeImages(cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeImagesInput{}

	result, err := client.DescribeImages(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil

}
