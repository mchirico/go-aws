package dynamoDB

import (
	"github.com/mchirico/go-aws/client"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

func (d *DB) Doc(location, aws string) *Doc {
	doc := &Doc{}
	doc.Location = location
	doc.AWS = aws
	return doc
}
