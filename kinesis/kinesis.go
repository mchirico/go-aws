package kinesis

/*
Ref:
https://docs.aws.amazon.com/streams/latest/dev/fundamental-stream.html
*/
import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

func Create(cfg aws.Config, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {

	client := kinesis.NewFromConfig(cfg)
	return client.CreateStream(context.TODO(), input)

}

func List(cfg aws.Config, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {

	client := kinesis.NewFromConfig(cfg)
	return client.ListStreams(context.TODO(), input)

}

func DescribeStream(cfg aws.Config, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {

	client := kinesis.NewFromConfig(cfg)
	return client.DescribeStream(context.TODO(), input)

}

func SubscribeToShard(cfg aws.Config, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
	client := kinesis.NewFromConfig(cfg)
	return client.SubscribeToShard(context.TODO(), input)
}

func Put(cfg aws.Config, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {

	client := kinesis.NewFromConfig(cfg)
	return client.PutRecord(context.TODO(), input)

}

func Get(cfg aws.Config, name string) (*kinesis.GetRecordsOutput, error) {

	shardId := "shardId-000000000000"
	input := &kinesis.GetShardIteratorInput{
		ShardId:           &shardId,
		ShardIteratorType: types.ShardIteratorTypeTrimHorizon,
		StreamName:        &name,
	}
	client := kinesis.NewFromConfig(cfg)
	result, err := client.GetShardIterator(context.TODO(), input)
	if err != nil {
		return nil, nil
	}
	input2 := &kinesis.GetRecordsInput{
		ShardIterator: result.ShardIterator,
	}
	return client.GetRecords(context.TODO(), input2)

}

func GetIter(cfg aws.Config, name string) (*kinesis.GetShardIteratorOutput, error) {

	shardId := "shardId-000000000000"
	input := &kinesis.GetShardIteratorInput{
		ShardId:           &shardId,
		ShardIteratorType: types.ShardIteratorTypeTrimHorizon,
		StreamName:        &name,
	}
	client := kinesis.NewFromConfig(cfg)
	return client.GetShardIterator(context.TODO(), input)

}

func GetIterSeq(cfg aws.Config, name, seq string) (*kinesis.GetShardIteratorOutput, error) {

	shardId := "shardId-000000000000"
	input := &kinesis.GetShardIteratorInput{
		ShardId:                &shardId,
		StartingSequenceNumber: &seq,
		ShardIteratorType:      types.ShardIteratorTypeAtSequenceNumber,
		StreamName:             &name,
	}
	client := kinesis.NewFromConfig(cfg)
	return client.GetShardIterator(context.TODO(), input)

}

func GetRecords(cfg aws.Config, shardIterator *string) (*kinesis.GetRecordsOutput, error) {
	client := kinesis.NewFromConfig(cfg)

	input := &kinesis.GetRecordsInput{
		ShardIterator: shardIterator,
	}
	return client.GetRecords(context.TODO(), input)

}

func GetStreamSeq(cfg aws.Config, name, seq string, num ...int32) (*kinesis.GetRecordsOutput, error) {
	client := kinesis.NewFromConfig(cfg)

	var limit int32 = 200
	if len(num) != 0 {
		limit = num[0]
	}
	result, err := GetIterSeq(cfg, name, seq)
	if err != nil {
		return nil, err
	}
	input := &kinesis.GetRecordsInput{
		ShardIterator: result.ShardIterator,
		Limit:         &limit,
	}
	return client.GetRecords(context.TODO(), input)

}

func Delete(cfg aws.Config, name string) (*kinesis.DeleteStreamOutput, error) {
	client := kinesis.NewFromConfig(cfg)

	enforceConsumerDeletion := true
	input := &kinesis.DeleteStreamInput{
		StreamName:              &name,
		EnforceConsumerDeletion: &enforceConsumerDeletion,
	}
	return client.DeleteStream(context.TODO(), input)

}

func Register(cfg aws.Config, consumerName, streamARN string) (*kinesis.RegisterStreamConsumerOutput, error) {

	input := &kinesis.RegisterStreamConsumerInput{
		ConsumerName: &consumerName,
		StreamARN:    &streamARN,
	}

	client := kinesis.NewFromConfig(cfg)
	return client.RegisterStreamConsumer(context.TODO(), input)

}
