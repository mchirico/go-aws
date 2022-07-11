package kinesis

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/mchirico/go-aws/client"
)

func TestCreate(t *testing.T) {
	var shards int32 = 1
	name := "mmc"
	input := &kinesis.CreateStreamInput{
		StreamName: &name,
		ShardCount: &shards,
	}
	result, err := Create(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestList(t *testing.T) {
	var limit int32 = 10
	input := &kinesis.ListStreamsInput{
		Limit: &limit,
	}
	result, err := List(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.StreamNames {
		fmt.Println(v)
	}

}

func put(input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	return Put(client.Config(), input)
}

func createData(name, key, seq string, data []byte) *kinesis.PutRecordInput {

	return &kinesis.PutRecordInput{
		Data:                      data,
		PartitionKey:              &key,
		StreamName:                &name,
		SequenceNumberForOrdering: &seq,
	}
}

func TestPut(t *testing.T) {

	put(createData("mmc", "key1", "1", []byte("one ...")))
	put(createData("mmc", "key2", "2", []byte("two ...")))

	result, err := put(createData("mmc", "key3", "3", []byte("last")))

	if err != nil {
		t.Fatal(err)
	}
	for _, v := range *result.SequenceNumber {
		fmt.Println(v)
	}

}

func TestGet(t *testing.T) {

	name := "mmc"
	result, err := Get(client.Config(), name)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range result.Records {
		fmt.Println(string(v.Data))
		fmt.Println(v.ApproximateArrivalTimestamp)

	}

}

func TestGetSeq(t *testing.T) {

	name := "mmc"
	seq := "2"
	result, err := GetStreamSeq(client.Config(), name, seq)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestDelete(t *testing.T) {
	name := "mmc"
	result, err := Delete(client.Config(), name)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
