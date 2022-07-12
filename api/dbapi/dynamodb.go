package dbapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/VladRomanciuc/Go-classes/api/models"
)
var tableName = "posts"

type dynamoDBTable struct {}

func NewDynamoDB() models.DbOps {
	return &dynamoDBTable{}
}

func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (table *dynamoDBTable) AddPost(post *models.Post) (*models.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Transforms the post to map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return nil, err
	}

	return post, err
}

func (table *dynamoDBTable) GetAll() ([]models.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}
	var posts []models.Post = []models.Post{}
	for _, i := range result.Items {
		post := models.Post{}

		err = dynamodbattribute.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (table *dynamoDBTable) GetById(id string) (*models.Post, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	post := models.Post{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	return &post, nil
}

// Delete: TODO
func (table *dynamoDBTable) Delete(post *models.Post) error {
	return nil
}

