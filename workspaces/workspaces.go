package workspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	ws "github.com/aws/aws-sdk-go-v2/service/workspaces"
	// "github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

// aws workspaces describe-workspaces
func DescribeWorkspaces(cfg aws.Config) (*ws.DescribeWorkspacesOutput, error) {
	client := ws.NewFromConfig(cfg)
	input := &ws.DescribeWorkspacesInput{}
	return client.DescribeWorkspaces(context.TODO(), input)

}
