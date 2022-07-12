package lambda

import (
	//"context"

	"context"
	"fmt"

	"archive/zip"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"os"
)

func CreateZip(file string) {
	// Create a buffer to write our archive to.

	os.Remove(file + ".zip")
	newZipFile, err := os.Create(file + ".zip")
	if err != nil {
		return
	}
	defer newZipFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(newZipFile)

	fs, err := os.Stat(file)
	if err != nil {
		return
	}
	b := make([]byte, fs.Size())
	f, err := os.Open(file)
	n, err := f.Read(b)
	if err != nil {
		fmt.Println("err read")
	}
	fmt.Println("Read: ", n)
	z, err := w.Create(file)
	if err != nil {
		return
	}

	n, err = z.Write(b)
	if err != nil {
		return
	}
	fmt.Println(n)
	err = w.Close()
	if err != nil {
		return
	}
	f.Close()

}

// GOOS=linux GOARCH=amd64  go build -o bin/lambda-time .
func Create(cfg aws.Config, name, handler, zipfile string,
	memorySize int32,
	timeOut int32,
	role string,
) error {
	client := lambda.NewFromConfig(cfg)
	// snippet-start:[lambda.go.create_function.read_zip]

	contents, err := ioutil.ReadFile(zipfile)
	if err != nil {
		fmt.Println("Got error trying to read " + zipfile)
		return err
	}

	createCode := &types.FunctionCode{
		//      S3Bucket:        bucket,
		//      S3Key:           zipFile,
		//      S3ObjectVersion: aws.String("1"),
		ZipFile: contents,
	}

	createArgs := &lambda.CreateFunctionInput{
		Code:         createCode,
		FunctionName: &name,
		Role:         &role,
		Handler:      &handler,
		MemorySize:   &memorySize,
		Runtime:      "go1.x",
		Timeout:      &timeOut,
	}

	result, err := client.CreateFunction(context.TODO(), createArgs)
	if err != nil {
		return err
	}
	_ = result

	return nil
}

func Delete(cfg aws.Config, name string) error {
	client := lambda.NewFromConfig(cfg)
	// snippet-start:[lambda.go.create_function.read_zip]

	result, err := client.DeleteFunction(context.TODO(), &lambda.DeleteFunctionInput{
		FunctionName: &name,
	})

	if err != nil {
		return err
	}
	_ = result

	return nil
}

// GOOS=linux GOARCH=amd64  go build -o bin/lambda-time .
func Invoke(cfg aws.Config, name, json string) (string, error) {
	client := lambda.NewFromConfig(cfg)

	result, err := client.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: &name,
		Payload:      []byte(json),
	})
	if err != nil {
		return "", err
	}

	return string(result.Payload), nil

}

func ListFunctions(cfg aws.Config, max int32) (*lambda.ListFunctionsOutput, error) {
	client := lambda.NewFromConfig(cfg)
	input := &lambda.ListFunctionsInput{
		MaxItems: &max,
	}

	result, err := client.ListFunctions(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, err

}

func GetFunctionConfiguration(cfg aws.Config, functionName string) (*lambda.GetFunctionConfigurationOutput, error) {
	client := lambda.NewFromConfig(cfg)
	input := &lambda.GetFunctionConfigurationInput{
		FunctionName: &functionName,
	}

	return client.GetFunctionConfiguration(context.TODO(), input)

}
