package workspaces

import (
	"github.com/mchirico/go-aws/client"
	"testing"
)

func TestDescribeWorkspaces(t *testing.T) {

	result, err := DescribeWorkspaces(client.Config())
	t.Log(result, err)

}
