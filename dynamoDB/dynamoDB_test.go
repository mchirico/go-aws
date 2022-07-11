package dynamoDB

import (
	"fmt"
	"github.com/mchirico/go-aws/client"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestCreateMovies(t *testing.T) {
	CreateMovie(client.Config(), "mmcMovies")
}

func TestCreatePKSK(t *testing.T) {
	Create(client.Config(), "mmcPKSK")
}

func TestPut_mmcPKSK(t *testing.T) {

	d := &Doc{
		Location: "location",
		AWS:      "bucket",
	}
	p := &PKSK{}
	p.PK = "My Data"
	p.SK = "Something"
	p.Status = "Good"
	p.Doc = *d

	av, err := attributevalue.MarshalMap(p)
	if err != nil {
		t.Fatal(err)
	}
	err = Put(client.Config(), "mmcPKSK", av)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_Get(t *testing.T) {
	table := "mmcPKSK"
	type KEY struct {
		PK string `json:"PK"`
		SK string `json:"SK"`
	}

	key, _ := attributevalue.MarshalMap(&KEY{
		PK: "My Data",
		SK: "Something",
	})

	input := &dynamodb.GetItemInput{
		Key:             key,
		TableName:       &table,
		AttributesToGet: []string{"PK", "Doc", "SK", "Status"},
	}
	result, err := Get(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	p := &PKSK{}
	attributevalue.UnmarshalMap(result.Item, p)
	if p.Status != "Good" {
		t.FailNow()
	}

}

func TestPut(t *testing.T) {
	items := getItems()

	for _, item := range items {
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			t.Fatal(err)
		}
		err = Put(client.Config(), "mmcMovies", av)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func GSI() []types.GlobalSecondaryIndexUpdate {
	gsi := []types.GlobalSecondaryIndexUpdate{
		{
			Create: &types.CreateGlobalSecondaryIndexAction{
				IndexName: aws.String("Tag"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Tag"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
	}

	return gsi
}

func TestGlobalIndex(t *testing.T) {
	tableName := "mmcMovies"
	input := &dynamodb.UpdateTableInput{
		TableName: &tableName,
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("Tag"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		GlobalSecondaryIndexUpdates: GSI(),
	}
	result, err := Modify(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestGlobalIndexPKSK(t *testing.T) {
	tableName := "mmcPKSK"
	input := &dynamodb.UpdateTableInput{
		TableName: &tableName,
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("Tag"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		GlobalSecondaryIndexUpdates: GSI(),
	}
	result, err := Modify(client.Config(), input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}

func TestDelete(t *testing.T) {
	err := Delete(client.Config(), "mmcMovies")
	if err != nil {
		t.Log(err)
	}
}
