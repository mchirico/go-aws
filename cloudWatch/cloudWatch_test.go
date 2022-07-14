package cloudWatch

import (
	"testing"

	"github.com/mchirico/go-aws/client"
)

func TestList(t *testing.T) {
	List(client.Config())
}

func Test_Logs(t *testing.T) {
	Logs(client.Config(), "prog2")
}
