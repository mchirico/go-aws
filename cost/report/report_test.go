package report

import (
	"fmt"
	"testing"

	"github.com/mchirico/go-aws/client"
)

func TestNewCdata(t *testing.T) {
	c := NewCdata(client.Config())
	c.cost()
}

func TestCdata_Max(t *testing.T) {
	c := NewCdata(client.Config())
	result := c.Max()
	fmt.Println(result)
}
