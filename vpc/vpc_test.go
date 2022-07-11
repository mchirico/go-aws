package vpc

import (
	"fmt"
	"github.com/mchirico/go-aws/client"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func TestListVPC(t *testing.T) {

	result, err := ListVPC(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.Vpcs {
		fmt.Println(*v.VpcId)
	}
}

func TestCreateEC2(t *testing.T) {

	keyPair := ""
	result, err := ListKeyPair(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result {
		keyPair = v
	}
	imageID := "ami-0cff7528ff583bf9a"
	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	CreateEC2(client.Config(), "types.InstanceTypeT2Micro", imageID, keyPair, subnetID)

}

func TestCreateEC2_on_Private(t *testing.T) {

	keyPair := ""
	result, err := ListKeyPair(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result {
		keyPair = v
	}
	imageID := "ami-0cff7528ff583bf9a"
	subnetID, _, err := GetSubnetID(client.Config(), "Private")
	CreateEC2(client.Config(), "types.InstanceTypeT2Micro", imageID, keyPair, subnetID)

}

func TestListEC2(t *testing.T) {
	result, err := ListEC2(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestListEC2ENI(t *testing.T) {
	result, instanceID, err := ListEC2ENI(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result, instanceID)
}

func GetDefaultAndSSH() ([]string, error) {
	out := []string{}
	name, groupID, err := ListSecurityGroups(client.Config())
	if err != nil {
		return out, err
	}
	for k, v := range name {
		if v == "default" {
			out = append(out, groupID[k])
		}
		if v == "ssh" {
			out = append(out, groupID[k])
		}
	}
	return out, nil
}

func Test_ModifyNetwork(t *testing.T) {
	result, instanceID, err := ListEC2ENI(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	networkInterfaceID := ""
	for _, v := range result {
		networkInterfaceID = v[0]
	}
	fmt.Println(result, instanceID, networkInterfaceID)

	groups, err := GetDefaultAndSSH()
	if err != nil {
		t.FailNow()
	}

	input := &ec2.ModifyNetworkInterfaceAttributeInput{
		NetworkInterfaceId: &networkInterfaceID,
		Groups:             groups,
		SourceDestCheck:    &types.AttributeBooleanValue{},
	}
	ModifyNetworkInterface(client.Config(), input)
}

func TestListSecurity(t *testing.T) {
	name, groupID, err := ListSecurityGroups(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(name, groupID)
}

func TestListKeyPair(t *testing.T) {

	result, err := ListKeyPair(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
func TestListSubnet(t *testing.T) {
	var maxResult int32 = 10
	input := &ec2.DescribeSubnetsInput{
		MaxResults: &maxResult,
	}
	result, err := ListSubnet(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.Subnets {
		fmt.Println(*v.SubnetId)
	}
}

func TestGetSubnetID(t *testing.T) {
	subnetID, vpcID, err := GetSubnetID(client.Config(), "Public")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*subnetID, " ", *vpcID)

}

func TestCreateVPC(t *testing.T) {

	key := "Name"
	value := "bozo"
	tag := types.Tag{Key: &key, Value: &value}
	tags := types.TagSpecification{
		ResourceType: "vpc",
		Tags: []types.Tag{
			tag,
		},
	}

	cidrBlock := "10.0.0.0/16"
	input := &ec2.CreateVpcInput{

		CidrBlock:         &cidrBlock,
		TagSpecifications: []types.TagSpecification{tags},
	}

	result, err := CreateVPC(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range *result.Vpc.CidrBlock {
		fmt.Println(v)
	}
}

func TestAddCidrToVPC(t *testing.T) {
	cidrBlock := "10.1.0.0/16"
	vpcId, err := GetVPCid(client.Config(), "bozo")
	input := &ec2.AssociateVpcCidrBlockInput{
		VpcId:     vpcId,
		CidrBlock: &cidrBlock,
	}
	result, err := AddCidrBlock(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func Name(name string, rType types.ResourceType) types.TagSpecification {
	key := "Name"
	value := name
	tag := types.Tag{Key: &key, Value: &value}
	tags := types.TagSpecification{
		ResourceType: rType,
		Tags: []types.Tag{
			tag,
		},
	}
	return tags
}

func TestAddPublicSubnetToVPC(t *testing.T) {

	avZone := "us-east-1a"
	cidrBlock := "10.0.0.0/24"
	vpcId, err := GetVPCid(client.Config(), "bozo")
	input := &ec2.CreateSubnetInput{
		VpcId:             vpcId,
		AvailabilityZone:  &avZone,
		CidrBlock:         &cidrBlock,
		TagSpecifications: []types.TagSpecification{Name("Public", types.ResourceTypeSubnet)},
	}
	result, err := AddSubnet(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestAddPrivateSubnetToVPC(t *testing.T) {

	avZone := "us-east-1a"
	cidrBlock := "10.0.16.0/20"
	vpcId, err := GetVPCid(client.Config(), "bozo")
	input := &ec2.CreateSubnetInput{
		VpcId:             vpcId,
		AvailabilityZone:  &avZone,
		CidrBlock:         &cidrBlock,
		TagSpecifications: []types.TagSpecification{Name("Private", types.ResourceTypeSubnet)},
	}
	result, err := AddSubnet(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestModifySubnet(t *testing.T) {

	// avZone := "us-east-1a"
	// cidrBlock := "10.0.16.0/20"
	// vpcId, err := GetBozo("bozo")
	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	mapIP := true

	input := &ec2.ModifySubnetAttributeInput{
		SubnetId:                       subnetID,
		MapPublicIpOnLaunch:            &types.AttributeBooleanValue{Value: &mapIP},
		PrivateDnsHostnameTypeOnLaunch: "",
	}
	result, err := ModifySubnet(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestAddGW(t *testing.T) {

	input := &ec2.CreateInternetGatewayInput{
		DryRun:            new(bool),
		TagSpecifications: []types.TagSpecification{Name("IGW", types.ResourceTypeInternetGateway)},
	}
	result, err := CreateGW(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func GetGWname() (*string, *string, error) {
	list, err := ListGW(client.Config())
	if err != nil {
		return nil, nil, err
	}
	for _, v := range list.InternetGateways {
		for _, tag := range v.Tags {
			if *tag.Key == "Name" {
				return v.InternetGatewayId, tag.Value, nil
			}
		}
	}
	return nil, nil, err
}

func TestAttachGW(t *testing.T) {

	gatewayID, _, err := GetGWname()
	if err != nil {
		t.Fatal(err)
	}

	_, vpcID, err := GetSubnetID(client.Config(), "Public")
	input := &ec2.AttachInternetGatewayInput{
		InternetGatewayId: gatewayID,
		VpcId:             vpcID,
	}
	result, err := AttachGW(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestCreateRT(t *testing.T) {

	_, vpcID, err := GetSubnetID(client.Config(), "Public")
	input := &ec2.CreateRouteTableInput{
		VpcId:             vpcID,
		TagSpecifications: []types.TagSpecification{Name("RTPublic", types.ResourceTypeRouteTable)},
	}
	result, err := CreateRT(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestCreatePrivateRT(t *testing.T) {

	_, vpcID, err := GetSubnetID(client.Config(), "Private")
	input := &ec2.CreateRouteTableInput{
		VpcId:             vpcID,
		TagSpecifications: []types.TagSpecification{Name("RTPrivate", types.ResourceTypeRouteTable)},
	}
	result, err := CreateRT(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestAssociateRT(t *testing.T) {

	routeID, _, err := GetRTid(client.Config(), "RTPublic")
	if err != nil {
		t.Fatal(err)
	}

	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	input := &ec2.AssociateRouteTableInput{
		RouteTableId: routeID,
		SubnetId:     subnetID,
	}
	result, err := AssociateRT(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestAssociatePrivateRT(t *testing.T) {

	routeID, _, err := GetRTid(client.Config(), "RTPrivate")
	if err != nil {
		t.Fatal(err)
	}

	subnetID, _, err := GetSubnetID(client.Config(), "Private")
	input := &ec2.AssociateRouteTableInput{
		RouteTableId: routeID,
		SubnetId:     subnetID,
	}
	result, err := AssociateRT(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

// Here we assign it to the GW
func TestAddRouteToRT(t *testing.T) {

	gatewayID, _, err := GetGWname()
	if err != nil {
		t.Fatal(err)
	}

	routeID, _, err := GetRTid(client.Config(), "RTPublic")
	if err != nil {
		t.Fatal(err)
	}

	destinationCidrBlock := "0.0.0.0/0"

	//	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	input := &ec2.CreateRouteInput{
		RouteTableId:         routeID,
		DestinationCidrBlock: &destinationCidrBlock,
		GatewayId:            gatewayID,
	}
	result, err := CreateRoute(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestAddNatRouteToRT(t *testing.T) {

	natGWID, err := NatGWID()
	if err != nil {
		t.Fatal(err)
	}

	routeID, _, err := GetRTid(client.Config(), "RTPrivate")
	if err != nil {
		t.Fatal(err)
	}

	destinationCidrBlock := "0.0.0.0/0"

	//	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	input := &ec2.CreateRouteInput{
		RouteTableId:         routeID,
		DestinationCidrBlock: &destinationCidrBlock,
		NatGatewayId:         natGWID,
	}
	result, err := CreateRoute(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func IP() (*ec2.AllocateAddressOutput, error) {
	input := &ec2.AllocateAddressInput{
		TagSpecifications: []types.TagSpecification{Name("NATIP", types.ResourceTypeElasticIp)},
	}
	r, err := AllocateIP(client.Config(), input)
	return r, err
}

func TestIP(t *testing.T) {
	//IP()
	key := "tag-key"
	value := "tag-value"
	k := []types.Filter{
		{
			Name:   &key,
			Values: []string{"Name"},
		},
		{
			Name:   &value,
			Values: []string{"NATIP"},
		},
	}
	input := &ec2.DescribeAddressesInput{
		Filters: k,
	}
	result, _ := ListIP(client.Config(), input)
	fmt.Println(result)
}

func TestFindIP(t *testing.T) {
	result, err := FindIP(client.Config(), "Name", "NATIP")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(result.Addresses[0].AllocationId)
}

func TestCreateNatGW(t *testing.T) {
	subnetID, _, err := GetSubnetID(client.Config(), "Public")
	resultIP, err := FindIP(client.Config(), "Name", "NATIP")
	allocationId := resultIP.Addresses[0].AllocationId
	input := &ec2.CreateNatGatewayInput{
		SubnetId:          subnetID,
		AllocationId:      allocationId,
		ConnectivityType:  types.ConnectivityTypePublic,
		TagSpecifications: []types.TagSpecification{Name("NATPublic", types.ResourceTypeNatgateway)},
	}
	result, err := CreateNatGW(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func NatGWID() (*string, error) {
	input := &ec2.DescribeNatGatewaysInput{}
	result, err := ListNatGW(client.Config(), input)
	if err != nil {
		return nil, err
	}
	return result.NatGateways[0].NatGatewayId, nil
}

func TestListNatGWFunc(t *testing.T) {

	result, _ := NatGWID()
	fmt.Println(result)
}

func TestListNatGW(t *testing.T) {
	input := &ec2.DescribeNatGatewaysInput{}
	result, err := ListNatGW(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestDeleteNatGW(t *testing.T) {

}

func TestDeleteVPC(t *testing.T) {

	result, err := ListVPC(client.Config())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.Vpcs {
		fmt.Println(*v.VpcId)
		err := DeleteVPC(client.Config(), *v.VpcId, false)
		if err != nil {
			t.Fatal(err)
		}
	}

}
