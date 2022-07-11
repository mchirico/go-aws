package vpc

import (
	"context"

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2DescribeVpcEndpointConnectionsAPI interface {
	DescribeVpcEndpointConnections(ctx context.Context,
		params *ec2.DescribeVpcEndpointConnectionsInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
}

func GetConnectionInfo(c context.Context,
	api EC2DescribeVpcEndpointConnectionsAPI,
	input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	return api.DescribeVpcEndpointConnections(context.Background(), input)
}

// Ref https://aws.github.io/aws-sdk-go-v2/docs/code-examples/ec2/describevpcendpoints/
func ListVPCConnections(cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVpcEndpointConnectionsInput{}

	resp, err := GetConnectionInfo(context.Background(), client, input)
	if err != nil {
		return err
	}
	cons := len(resp.VpcEndpointConnections)

	if cons == 0 {
		fmt.Println("Could not find any VCP endpoint connections")

	}

	fmt.Println("VPC endpoint: Details:")
	respDecrypted, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(respDecrypted))

	fmt.Println()
	fmt.Println("Found " + strconv.Itoa(cons) + " VCP endpoint connection(s)")

	return nil

}

func CreateVPC(cfg aws.Config, input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateVpc(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func AddCidrBlock(cfg aws.Config, input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.AssociateVpcCidrBlock(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func AddSubnet(cfg aws.Config, input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateSubnet(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func ModifySubnet(cfg aws.Config, input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.ModifySubnetAttribute(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func ListSubnet(cfg aws.Config, input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.DescribeSubnets(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func CreateRT(cfg aws.Config, input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateRouteTable(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func AssociateRT(cfg aws.Config, input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.AssociateRouteTable(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func CreateRoute(cfg aws.Config, input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateRoute(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func CreateGW(cfg aws.Config, input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateInternetGateway(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func AttachGW(cfg aws.Config, input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.AttachInternetGateway(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func ListRT(cfg aws.Config) (*ec2.DescribeRouteTablesOutput, error) {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeRouteTablesInput{}
	resp, err := client.DescribeRouteTables(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func ListGW(cfg aws.Config) (*ec2.DescribeInternetGatewaysOutput, error) {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInternetGatewaysInput{}
	resp, err := client.DescribeInternetGateways(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func ListVPC(cfg aws.Config) (*ec2.DescribeVpcsOutput, error) {
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVpcsInput{}
	resp, err := client.DescribeVpcs(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func DeleteVPC(cfg aws.Config, vpcID string, dryrun bool) error {
	client := ec2.NewFromConfig(cfg)

	result, err := client.DeleteVpc(context.TODO(), &ec2.DeleteVpcInput{
		VpcId:  &vpcID,
		DryRun: &dryrun,
	})

	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil

}

func GetVPCid(cfg aws.Config, name string) (*string, error) {
	result, err := ListVPC(cfg)
	if err != nil {
		return nil, err
	}
	for _, v := range result.Vpcs {
		for _, tag := range v.Tags {
			if *tag.Value == name && *tag.Key == "Name" {
				return v.VpcId, nil
			}
		}

	}
	return nil, nil
}

func GetRTid(cfg aws.Config, name string) (*string, *string, error) {
	result, err := ListRT(cfg)
	if err != nil {
		return nil, nil, err
	}
	for _, v := range result.RouteTables {
		for _, tag := range v.Tags {
			if *tag.Value == name && *tag.Key == "Name" {
				return v.RouteTableId, v.VpcId, nil
			}
		}

	}
	return nil, nil, nil
}

// Returns subnetID,vpcID,error
func GetSubnetID(cfg aws.Config, name string) (*string, *string, error) {
	var maxResult int32 = 10
	input := &ec2.DescribeSubnetsInput{
		MaxResults: &maxResult,
	}

	result, err := ListSubnet(cfg, input)
	if err != nil {
		return nil, nil, err
	}
	for _, v := range result.Subnets {
		for _, tag := range v.Tags {
			if *tag.Value == name && *tag.Key == "Name" {
				return v.SubnetId, v.VpcId, nil
			}
		}

	}
	return nil, nil, nil
}

func CreateEC2(cfg aws.Config, instanceType, imageID string, keyPair string, subnetId *string) error {
	client := ec2.NewFromConfig(cfg)

	result, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
		InstanceType: types.InstanceTypeT2Micro,
		ImageId:      aws.String(imageID),
		KeyName:      &keyPair,
		SubnetId:     subnetId,
	})
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
func AllocateIP(cfg aws.Config, input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	client := ec2.NewFromConfig(cfg)
	result, err := client.AllocateAddress(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindIP(cfg aws.Config, name, value string) (*ec2.DescribeAddressesOutput, error) {
	client := ec2.NewFromConfig(cfg)

	key := "tag-key"
	tagValue := "tag-value"
	k := []types.Filter{
		{
			Name:   &key,
			Values: []string{name},
		},
		{
			Name:   &tagValue,
			Values: []string{value},
		},
	}
	input := &ec2.DescribeAddressesInput{
		Filters: k,
	}
	result, err := client.DescribeAddresses(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func ListIP(cfg aws.Config, input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	client := ec2.NewFromConfig(cfg)
	result, err := client.DescribeAddresses(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteIP(cfg aws.Config, input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.ReleaseAddress(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateNatGW(cfg aws.Config, input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.CreateNatGateway(context.TODO(), input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func ListNatGW(cfg aws.Config, input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.DescribeNatGateways(context.TODO(), input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func DeleteNatGW(cfg aws.Config, input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.DeleteNatGateway(context.TODO(), input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func ModifyNetworkInterface(cfg aws.Config, input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	client := ec2.NewFromConfig(cfg)

	result, err := client.ModifyNetworkInterfaceAttribute(context.TODO(), input)
	if err != nil {
		return result, err
	}
	fmt.Println(result)
	return result, nil
}

func ListKeyPair(cfg aws.Config) ([]string, error) {
	out := []string{}
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeKeyPairsInput{}
	resp, err := client.DescribeKeyPairs(context.TODO(), input)
	if err != nil {
		return out, err
	}
	for _, v := range resp.KeyPairs {
		out = append(out, *v.KeyName)
	}

	return out, nil
}

func ListEC2(cfg aws.Config) ([]string, error) {
	out := []string{}
	client := ec2.NewFromConfig(cfg)

	var max int32 = 10
	input := &ec2.DescribeInstancesInput{
		MaxResults: &max,
	}
	resp, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return out, err
	}
	for _, v := range resp.Reservations {
		for _, instances := range v.Instances {

			out = append(out, *instances.InstanceId)
		}

	}

	return out, nil
}

func ListEC2ENI(cfg aws.Config) (map[string][]string, []string, error) {
	out := []string{}
	m := map[string][]string{}
	client := ec2.NewFromConfig(cfg)

	var max int32 = 10
	input := &ec2.DescribeInstancesInput{
		MaxResults: &max,
	}
	resp, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return m, out, err
	}
	for _, v := range resp.Reservations {
		for _, instances := range v.Instances {
			if instances.State.Name != "running" {
				continue
			}

			out = append(out, *instances.InstanceId)
			inter := []string{}
			for _, nis := range instances.NetworkInterfaces {
				inter = append(inter, *nis.NetworkInterfaceId)

			}
			m[*instances.InstanceId] = inter

		}

	}

	return m, out, nil
}

func ListSecurityGroups(cfg aws.Config) ([]string, []string, error) {
	name := []string{}
	groupID := []string{}
	client := ec2.NewFromConfig(cfg)
	var max int32 = 100

	input := &ec2.DescribeSecurityGroupsInput{
		MaxResults: &max,
	}
	resp, err := client.DescribeSecurityGroups(context.TODO(), input)
	if err != nil {
		return name, groupID, err
	}
	for _, v := range resp.SecurityGroups {
		name = append(name, *v.GroupName)
		groupID = append(groupID, *v.GroupId)
	}

	return name, groupID, nil
}
