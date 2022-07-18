package dynamoDB

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"testing"
	"time"
)

func TestD_List(t *testing.T) {
	d := NewDB("pksk")
	result, err := d.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range result.TableNames {
		fmt.Println(table)
	}

}

func TestD_Put(t *testing.T) {
	pkey := "TestD_Put"
	skey := "skey:TestD_Put2"

	d := NewDB("pksk")
	p := &PKSK{}
	p.PK = pkey
	p.SK = skey
	p.Status = "TimeStamp: " + time.Now().Format(time.RFC3339)
	p.GSI = "GSI-must have value.. this is there"
	p.Doc = *d.Doc("name", time.Now().Format(time.RFC3339), "{key:value}")
	av, err := attributevalue.MarshalMap(p)

	if err != nil {
		t.Fatal(err)
	}
	_, err = Put(d.cfg, d.name, av)
	if err != nil {
		t.Fatal(err)
	}

	result, err := d.Get(pkey, skey)
	if err != nil {
		t.Fatal(err)
	}
	if result.PK != pkey || result.SK != skey {
		t.Fatal("Get failed")
	}
}

func Test_Update(t *testing.T) {
	pkey := "TestD_Put"
	skey := "skey:TestD_Put5"

	d := NewDB("pksk")

	doc := d.Doc("Pizza Fuzz...1 2 3 4 5", time.Now().Format(time.RFC3339), "{key:value}")
	av, err := attributevalue.MarshalMap(doc)

	m := map[string]types.AttributeValue{}
	m[":doc"] = &types.AttributeValueMemberM{Value: av}

	_, err = d.UpdateDoc(pkey, skey, m)
	if err != nil {
		t.Fatal(err)
	}

	result, err := d.Get(pkey, skey)
	if err != nil {
		t.Fatal(err)
	}
	if result.PK != pkey || result.SK != skey {
		t.Fatal("Get failed")
	}
}

func Test_DeleteItem(t *testing.T) {
	pkey := "TestD_Put"
	skey := "skey:TestD_Put5"

	d := NewDB("pksk")

	_, err := d.DeleteItem(pkey, skey)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_Query(t *testing.T) {

	d := NewDB("pksk")

	expAttValues := map[string]types.AttributeValue{}
	expAttValues[":name"] = &types.AttributeValueMemberS{Value: "GSI-search"}

	result, err := d.Query("GSI", "GSI = :name", expAttValues)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range result.Items {
		for k, v := range item {
			fmt.Println(k, v)
		}
	}
	fmt.Println(result.Items)

}
