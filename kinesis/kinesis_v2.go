package kinesis

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/mchirico/go-aws/client"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

type P struct {
	seq        int64
	client     aws.Config
	name       *string
	streamARN  *string
	shardCount *int32
}

func NewP(name string) *P {
	var shardCount int32 = 1
	return &P{client: client.Config(),
		name:       &name,
		shardCount: &shardCount,
	}
}

func (p *P) ShareCount(i int32) *int32 {
	p.shardCount = &i
	return p.shardCount
}

func (p *P) seqOrder() *string {
	seq := fmt.Sprintf("%d", p.seq)
	p.seq += 1
	return &seq
}

func (p *P) Name(name string) *P {
	p.name = &name
	return p
}

func (p *P) Create() (*kinesis.CreateStreamOutput, error) {
	input := &kinesis.CreateStreamInput{
		StreamName: p.name,
		ShardCount: p.shardCount,
	}
	return Create(p.client, input)

}

func (p *P) Get() (*kinesis.GetRecordsOutput, error) {
	return Get(p.client, *p.name)
}

func (p *P) List() (*kinesis.ListStreamsOutput, error) {
	var max int32 = 10
	input := &kinesis.ListStreamsInput{
		Limit: &max,
	}
	return List(p.client, input)
}

func (p *P) DescribeStream() (*kinesis.DescribeStreamOutput, error) {
	var max int32 = 10
	input := &kinesis.DescribeStreamInput{
		StreamName: p.name,
		Limit:      &max,
	}
	return DescribeStream(p.client, input)
}

func (p *P) StreamARN() (*string, error) {
	result, err := p.DescribeStream()
	if err != nil {
		return nil, err
	}
	return result.StreamDescription.StreamARN, nil
}

func (p *P) Delete() (*kinesis.DeleteStreamOutput, error) {
	return Delete(p.client, *p.name)
}

func (p *P) Put(key string, data []byte) (*kinesis.PutRecordOutput, error) {
	input := &kinesis.PutRecordInput{
		Data:                      data,
		PartitionKey:              &key,
		StreamName:                p.name,
		SequenceNumberForOrdering: p.seqOrder(),
	}
	return Put(p.client, input)
}

func (p *P) SubscribeToShard() (*kinesis.SubscribeToShardOutput, error) {
	shardId := "shardId-000000000000"
	var startPosition types.StartingPosition
	startPosition.Type = "TRIM_HORIZON"
	input := &kinesis.SubscribeToShardInput{
		ShardId:          &shardId,
		StartingPosition: &startPosition,
	}
	return SubscribeToShard(p.client, input)
}

func (p *P) Register(consumerName string, streamARN string) (*kinesis.RegisterStreamConsumerOutput, error) {

	return Register(p.client, consumerName, streamARN)
}
