package cost

/*
Old version of cost:
https://docs.aws.amazon.com/code-samples/latest/catalog/go-costexplorer-get_cost_and_usage.go.html
*/
import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func Cost(cfg aws.Config, startDate, endDate string) (*costexplorer.GetCostAndUsageOutput, error) {
	client := costexplorer.NewFromConfig(cfg)

	metrics := []string{
		"BlendedCost",
		"UnblendedCost",
		"UsageQuantity",
	}
	service := "SERVICE"
	groupDef := types.GroupDefinition{Key: &service, Type: types.GroupDefinitionTypeDimension}

	input := &costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityDaily,
		Metrics:     metrics,
		TimePeriod:  &types.DateInterval{End: &endDate, Start: &startDate},
		GroupBy: []types.GroupDefinition{
			groupDef,
		},
	}
	return client.GetCostAndUsage(context.TODO(), input)
}

func CostReport(cfg aws.Config, days int) {
	endDate := time.Now().Add(time.Duration(24) * time.Hour).Format("2006-01-02")
	startDate := time.Now().Add(time.Duration(-24*days) * time.Hour).Format("2006-01-02")
	result, err := Cost(cfg, startDate, endDate)
	if err != nil {
		return
	}
	var grandTotal float64
	for _, v := range result.ResultsByTime {
		fmt.Printf("\n\n%-20s  %-20s\n", *v.TimePeriod.Start, *v.TimePeriod.End)
		var total float64
		for _, g := range v.Groups {
			amount := *g.Metrics["BlendedCost"].Amount
			if amount, err := strconv.ParseFloat(amount, 32); err == nil {
				total = total + amount
			}
			if amount != "0" {
				fmt.Printf("%-20s  %-20s\n", *g.Metrics["BlendedCost"].Amount, g.Keys[0])
			}
		}
		grandTotal += total
		fmt.Printf("__________________________________\n%-20f  %-20s\n", total, "Total")
	}
	fmt.Printf("\n\n__________________________________\n%-20f  %-20s\n\n\n", grandTotal, "Grand Total")

}

func CostRecords(cfg aws.Config, days int) {
	endDate := time.Now().Add(time.Duration(24) * time.Hour).Format("2006-01-02")
	startDate := time.Now().Add(time.Duration(-24*days) * time.Hour).Format("2006-01-02")
	result, err := Cost(cfg, startDate, endDate)
	if err != nil {
		return
	}
	var grandTotal float64
	for _, v := range result.ResultsByTime {
		fmt.Printf("\n\n%-20s  %-20s\n", *v.TimePeriod.Start, *v.TimePeriod.End)
		var total float64
		for _, g := range v.Groups {
			amount := *g.Metrics["BlendedCost"].Amount
			if amount, err := strconv.ParseFloat(amount, 32); err == nil {
				total = total + amount
			}
			if amount != "0" {
				fmt.Printf("%-20s  %-20s\n", *g.Metrics["BlendedCost"].Amount, g.Keys[0])
			}
		}
		grandTotal += total
		fmt.Printf("__________________________________\n%-20f  %-20s\n", total, "Total")
	}
	fmt.Printf("\n\n__________________________________\n%-20f  %-20s\n\n\n", grandTotal, "Grand Total")

}
