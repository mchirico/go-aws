package dynamoDB

// snippet-start:[dynamodb.go.create_table.imports]
import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Doc struct {
	Name      string `json:"Name"`
	Timestamp string `json:"Timestamp"`
	JSON      string `json:"JSON"`
}

type PKSK struct {
	PK     string `json:"PK"`
	SK     string `json:"SK"`
	Status string `json:"Status"`
	GSI    string `json:"GSI"`
	Doc    Doc    `json:"Doc"`
}

type Item struct {
	Year   int     `json:"Year"`
	Title  string  `json:"Title"`
	Plot   string  `json:"Plot"`
	Rating float64 `json:"Rating"`
	Status string  `json:"Status"`
}

// Get table items from JSON file
func getItems() []Item {

	resp, err := http.Get("https://storage.googleapis.com/montco-stats/movieStatus.json") //use package "net/http"

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}

	var items []Item
	json.Unmarshal(body, &items)
	fmt.Printf("items: %v\n", items[0].Year)
	fmt.Printf("items: %v\n", items[1].Year)
	return items
}

func Delete(cfg aws.Config, tableName string) error {

	client := dynamodb.NewFromConfig(cfg)
	result, err := client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: &tableName})
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil

}

func Put(cfg aws.Config, tableName string, av map[string]types.AttributeValue) (*dynamodb.PutItemOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	return client.PutItem(context.TODO(), input)

}

func Get(cfg aws.Config, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	client := dynamodb.NewFromConfig(cfg)

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Query(cfg aws.Config, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	client := dynamodb.NewFromConfig(cfg)

	result, err := client.Query(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func UpdateTable(cfg aws.Config, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	result, err := client.UpdateTable(context.TODO(), input)
	return result, err
}

func UpdateItem(cfg aws.Config, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	result, err := client.UpdateItem(context.TODO(), input)
	return result, err
}

func DeleteItem(cfg aws.Config, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	result, err := client.DeleteItem(context.TODO(), input)
	return result, err
}

func List(cfg aws.Config, input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	result, err := client.ListTables(context.TODO(), input)
	return result, err
}

func CreateMovie(cfg aws.Config, tableName string) {

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("Status"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       types.KeyTypeRange,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("Status"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Status"),
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
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},

		TableName: aws.String(tableName),
	}

	_, err := client.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)

}

func Create(cfg aws.Config, tableName string) {

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("Status"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("GSI"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("Status"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Status"),
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
			{
				IndexName: aws.String("GSI"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("GSI"),
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
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},

		TableName: aws.String(tableName),
	}

	_, err := client.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)

}
