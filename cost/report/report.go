package report

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mchirico/go-aws/cost"
)

type rec struct {
	startDay string
	endDay   string
	cost     float64
	service  string
}

type Cdata struct {
	cfg  aws.Config
	data []rec
}

func NewCdata(cfg aws.Config) *Cdata {
	return &Cdata{cfg: cfg, data: []rec{}}
}

func (c *Cdata) Max() *rec {
	c.cost()
	r := &rec{}
	var max float64
	for _, v := range c.data {
		if v.cost > max {
			max = v.cost
			r.cost = v.cost
			r.service = v.service
			r.startDay = v.startDay
			r.endDay = v.endDay
		}
	}
	return r
}

func (c *Cdata) cost() {
	endDate := time.Now().Add(time.Duration(24) * time.Hour).Format("2006-01-02")
	startDate := time.Now().Add(time.Duration(-31*24) * time.Hour).Format("2006-01-02")
	result, err := cost.Cost(c.cfg, startDate, endDate)
	if err != nil {
		return
	}

	for _, v := range result.ResultsByTime {
		r := rec{}
		r.startDay = *v.TimePeriod.Start
		r.endDay = *v.TimePeriod.End
		for _, g := range v.Groups {
			amount := *g.Metrics["BlendedCost"].Amount
			if amount == "0" {
				continue
			}
			if amount, err := strconv.ParseFloat(amount, 32); err == nil {
				r.cost = amount
				r.service = g.Keys[0]
			}

		}
		c.data = append(c.data, r)
	}
}
