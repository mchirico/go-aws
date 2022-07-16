package dynamoDB

import (
	"github.com/mchirico/go-aws/client"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type D struct {
	name string
	cfg  aws.Config
	max  int32
}

func NewD() *D {
	return &D{cfg: client.Config(), max: int32(50)}
}

func (d *D) List() (*dynamodb.ListTablesOutput, error) {
	input := &dynamodb.ListTablesInput{
		Limit: &d.max,
	}
	return List(d.cfg, input)

}
