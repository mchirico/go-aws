package dynamoDB

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"testing"
	"time"
)

func TestD_List(t *testing.T) {
	d := NewDB("mmcPKSK")
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
	skey := "skey:TestD_Put"

	d := NewDB("mmcPKSK")
	p := &PKSK{}
	p.PK = pkey
	p.SK = skey
	p.Status = "Good"
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
