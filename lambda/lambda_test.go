package lambda

import (
	"fmt"

	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/mchirico/go-aws/client"
	"github.com/mchirico/go-aws/iam"
)

func TestCreateZip(t *testing.T) {
	CreateZip("main")
}

func TestCreate(t *testing.T) {

	i := iam.NewI("k8s")
	role, err := i.GetRole("Test-Role")
	if err != nil {
		t.Fatal(err)
	}

	err = Create(client.Config(), "prog2", "main", "main.zip",
		128,
		10, *role)

	if err != nil {
		t.Fatal(err)
	}

}

func TestCreateSNS(t *testing.T) {

	i := iam.NewI("appleM1")
	role, err := i.GetRole("Test-Role")
	if err != nil {
		t.Fatal(err)
	}

	err = Create(client.Config(), "sns", "main", "main.zip",
		128,
		10, *role)

	if err != nil {
		t.Fatal(err)
	}

}

func TestDelete(t *testing.T) {
	err := Delete(client.Config(), "sns")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ListFunctions(t *testing.T) {
	result, err := ListFunctions(client.Config(), 30)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_ListEvents(t *testing.T) {
	result, err := ListEvents(client.Config(), "sns")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func Test_GetFunctionEventInvokeConfig(t *testing.T) {

	result, err := GetFunctionEventInvokeConfig(client.Config(), "sns")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestFuncConfig(t *testing.T) {
	result, err := GetFunctionConfiguration(client.Config(), "sns")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestInvoke(t *testing.T) {

	json := `{
		"name": "mike",
		"age": 120
	  }
	`
	result, err := Invoke(client.Config(), "prog2", json)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func TestInvoke2(t *testing.T) {

	result, err := Invoke(client.Config(), "prog2", KinesisExample())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}

func Test_EventSourceMapping(t *testing.T) {

	var max int32 = 10
	input := &lambda.ListEventSourceMappingsInput{
		MaxItems: &max,
	}

	result, err := ListEventSourceMapping(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result.EventSourceMappings)

	for _, v := range result.EventSourceMappings {
		fmt.Println(*v.FunctionArn)
	}
}

func Test_FindARN(t *testing.T) {
	result, err := FindARN(client.Config(), "sns")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
