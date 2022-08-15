package cloudWatch

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

// CWListMetricsAPI defines the interface for the ListMetrics function.
// We use this interface to test the function using a mocked service.
type CWListMetricsAPI interface {
	ListMetrics(ctx context.Context,
		params *cloudwatch.ListMetricsInput,
		optFns ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricsOutput, error)
}

// GetMetrics gets the name, namespace, and dimension name of your Amazon CloudWatch metrics
// Inputs:
//
//	c is the context of the method call, which includes the Region
//	api is the interface that defines the method call
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a ListMetricsOutput object containing the result of the service call and nil
//	Otherwise, nil and an error from the call to ListMetrics
func GetMetrics(c context.Context, api CWListMetricsAPI, input *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	return api.ListMetrics(c, input)
}

func List(cfg aws.Config) {

	client := cloudwatch.NewFromConfig(cfg)

	input := &cloudwatch.ListMetricsInput{}

	result, err := GetMetrics(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Could not get metrics")
		return
	}

	fmt.Println("Metrics:")
	numMetrics := 0

	for _, m := range result.Metrics {
		fmt.Println("   Metric Name: " + *m.MetricName)
		fmt.Println("   Namespace:   " + *m.Namespace)
		fmt.Println("   Dimensions:")
		for _, d := range m.Dimensions {
			fmt.Println("      " + *d.Name + ": " + *d.Value)
		}

		fmt.Println("")
		numMetrics++
	}

	fmt.Println("Found " + strconv.Itoa(numMetrics) + " metrics")
}

// https://www.youtube.com/watch?v=aZ-gP4rbFDo
//aws logs tail "/aws/lambda/sns" --follow

func Logs(cfg aws.Config, lgroups string) {

	client := cloudwatchlogs.NewFromConfig(cfg)

	prefix := "/aws/lambda/"
	var max int32 = 20
	dinput := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit:              &max,
		LogGroupNamePrefix: &prefix,
	}

	result, err := client.DescribeLogGroups(context.TODO(), dinput)
	if err != nil {
		return
	}

	_ = result

	sinput := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &lgroups,
		Limit:        &max,
	}
	r, err := client.DescribeLogStreams(context.TODO(), sinput)
	if err != nil {
		return
	}
	fmt.Println(r)
	if len(r.LogStreams) == 0 {
		fmt.Println("No LogStreams found")
		return
	}
	logStream := r.LogStreams[len(r.LogStreams)-1].LogStreamName

	input := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &lgroups,
		LogStreamName: logStream,
	}
	GetLogEvents(cfg, input)

}

func GetLogEvents(cfg aws.Config, input *cloudwatchlogs.GetLogEventsInput) {
	client := cloudwatchlogs.NewFromConfig(cfg)
	result, err := client.GetLogEvents(context.TODO(), input)
	if err != nil {
		return
	}
	for _, e := range result.Events {
		fmt.Println(*e.Message)
	}

}

func DeleteLogStream(cfg aws.Config, input *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
	client := cloudwatchlogs.NewFromConfig(cfg)
	return client.DeleteLogStream(context.TODO(), input)

}
