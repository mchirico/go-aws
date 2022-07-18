package dynamoDB

import (
	"github.com/mchirico/go-aws/client"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DB struct {
	name string
	cfg  aws.Config
	max  int32
}

func NewDB(name string) *DB {
	return &DB{
		name: name,
		cfg:  client.Config(),
		max:  int32(50)}
}

func (d *DB) List() (*dynamodb.ListTablesOutput, error) {
	input := &dynamodb.ListTablesInput{
		Limit: &d.max,
	}
	return List(d.cfg, input)

}

/*
https://dynobase.dev/dynamodb-golang-query-examples/#query-index

svc := dynamodb.NewFromConfig(cfg)
    out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
        TableName:              aws.String("my-table"),
        IndexName:              aws.String("GSI1"),
        KeyConditionExpression: aws.String("gsi1pk = :gsi1pk and gsi1sk > :gsi1sk"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":gsi1pk": &types.AttributeValueMemberS{Value: "123"},
            ":gsi1sk": &types.AttributeValueMemberN{Value: "20150101"},
        },
    })

aws dynamodb query \
 --table-name PKSK \
 --index-name GSI \
 --key-condition-expression "GSI = :name" \
 --expression-attribute-values '{":name":{"S":"GSI-search"}}'

*/
func (d *DB) Query(index, keyConditionExpression string, expAttValues map[string]types.AttributeValue) (*dynamodb.QueryOutput, error) {
	input := &dynamodb.QueryInput{
		TableName:                 &d.name,
		IndexName:                 &index,
		KeyConditionExpression:    &keyConditionExpression,
		ExpressionAttributeValues: expAttValues,
	}

	return Query(d.cfg, input)

}

func (d *DB) Put(pkey, skey, status string, doc *Doc) error {

	p := &PKSK{}
	p.PK = pkey
	p.SK = skey
	p.Status = status
	p.Doc = *doc

	av, err := attributevalue.MarshalMap(p)
	if err != nil {
		return err
	}
	_, err = Put(d.cfg, d.name, av)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateDoc(pkey, skey string, av map[string]types.AttributeValue) (*dynamodb.UpdateItemOutput, error) {

	input := &dynamodb.UpdateItemInput{
		TableName: &d.name,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkey},
			"SK": &types.AttributeValueMemberS{Value: skey},
		},
		UpdateExpression:          aws.String("set Doc = :doc"),
		ExpressionAttributeValues: av,
	}
	return UpdateItem(d.cfg, input)

}

func (d *DB) Get(pkey, skey string) (*PKSK, error) {

	type KEY struct {
		PK string `json:"PK"`
		SK string `json:"SK"`
	}

	key, _ := attributevalue.MarshalMap(&KEY{
		PK: pkey,
		SK: skey,
	})

	input := &dynamodb.GetItemInput{
		Key:             key,
		TableName:       &d.name,
		AttributesToGet: []string{"PK", "Doc", "SK", "Status"},
	}
	result, err := Get(d.cfg, input)
	if err != nil {
		return nil, err
	}
	p := &PKSK{}
	err = attributevalue.UnmarshalMap(result.Item, p)
	if err != nil {
		return nil, err
	}
	return p, nil

}

func (d *DB) Doc(name, timestamp, json string) *Doc {
	doc := &Doc{
		Name:      name,
		Timestamp: timestamp,
		JSON:      json,
	}

	return doc
}
