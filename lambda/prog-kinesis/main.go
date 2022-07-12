package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"prog/kinesis"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	kinesis.Put()
	return MyResponse{Message: fmt.Sprintf("v2 %s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
