package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"prog/kinesis"

    "github.com/aws/aws-lambda-go/events"
)



func handler(ctx context.Context, kinesisEvent events.KinesisEvent) {
	
    for _, record := range kinesisEvent.Records {
        kinesisRecord := record.Kinesis
        dataBytes := kinesisRecord.Data
        dataText := string(dataBytes)
		kinesis.Put(dataBytes)
        fmt.Printf("%s Data = %s \n", record.EventName, dataText)
    }
}


func main() {
	lambda.Start(handler)
}
