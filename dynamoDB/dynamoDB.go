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
	Location string `json:"PK"`
	AWS      string `json:"PK"`
}

type PKSK struct {
	PK     string `json:"PK"`
	SK     string `json:"SK"`
	Status string `json:"Status"`
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

func Put(cfg aws.Config, tableName string, av map[string]types.AttributeValue) error {
	client := dynamodb.NewFromConfig(cfg)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	result, err := client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func Get(cfg aws.Config, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	client := dynamodb.NewFromConfig(cfg)

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Modify(cfg aws.Config, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	client := dynamodb.NewFromConfig(cfg)
	result, err := client.UpdateTable(context.TODO(), input)
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
