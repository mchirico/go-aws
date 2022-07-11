package vpc

import (
	"github.com/mchirico/go-aws/client"
	"testing"
)

func TestDescribeTypeOfferings(t *testing.T) {
	DescribeTypeOfferings(client.Config())
}

func TestDescribeImages(t *testing.T) {
	DescribeImages(client.Config())
}
