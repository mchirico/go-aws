package lambda

import (
	"fmt"
	"github.com/mchirico/go-aws/client"
	"github.com/mchirico/go-aws/iam"
	"testing"
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

	Create(client.Config(), "prog2", "main", "main.zip",
		128,
		10, *role)
	json := `{
		"name": "value1",
		"age": 59
	  }
	`
	Invoke(client.Config(), "prog2", json)

}

func TestDelete(t *testing.T) {
	Delete(client.Config(), "prog2")
}

func TestInvoke(t *testing.T) {

	json := `{
		"name": "mike",
		"age": 60
	  }
	`
	result, err := Invoke(client.Config(), "prog2", json)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(result)

}