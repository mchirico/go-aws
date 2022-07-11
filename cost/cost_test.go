package cost

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/mchirico/go-aws/client"
)

func TestCost(t *testing.T) {

	endDate := time.Now().Add(time.Duration(24) * time.Hour).Format("2006-01-02")
	startDate := time.Now().Add(time.Duration(-24*10) * time.Hour).Format("2006-01-02")
	result, err := Cost(client.Config(), startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	var gtotal float64
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
		gtotal += total
		fmt.Printf("__________________________________\n%-20f  %-20s\n", total, "Total")
	}
	fmt.Printf("\n\n__________________________________\n%-20f  %-20s\n\n\n", gtotal, "Grand Total")
}
