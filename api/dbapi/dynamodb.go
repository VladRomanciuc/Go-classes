package dbapi

import (

	"log"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/spf13/viper"
)

var tableName = "posts"

type dynamoDBTable struct {
	tableName string
}

func getEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
	return value
}

func NewDynamoDB() models.DbOps {
	return &dynamoDBTable{
		tableName: tableName,
	}
}

func createDynamoDBClient() *dynamodb.Client {
	c := context.Background()
	cfg, err := config.LoadDefaultConfig(c,
		config.WithRegion(getEnv("Region")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	awsClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(getEnv("KeyID"), getEnv("Key"), "")
	})
	return awsClient
}

func (table *dynamoDBTable) AddPost(post *models.Post) (*models.Post, error) {
	c := context.Background()
	// Get a new DynamoDB client
	
	dynamoDBClient := createDynamoDBClient()

	// Transforms the post to map[string]*dynamodb.AttributeValue
	_, err := dynamoDBClient.PutItem(c, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: post.Id},
			"title": &types.AttributeValueMemberS{Value: post.Title},
			"text": &types.AttributeValueMemberS{Value: post.Text},
			},	
	})
	if err != nil {
		return nil, err
	}
	return post, err
}

func (table *dynamoDBTable) GetAll() ([]models.Post, error) {
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(c, params)
	if err != nil {
		return nil, err
	}

	var posts []models.Post = []models.Post{}
	for _, i := range result.Items {
		post := models.Post{}

		err = attributevalue.UnmarshalMap(i, &post)

		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (table *dynamoDBTable) GetById(id string) (*models.Post, error) {
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.GetItem(c, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *aws.String(id)},
		},
	})
	if err != nil {
		return nil, err
	}
	post := models.Post{}
	err = attributevalue.UnmarshalMap(result.Item, &post)
	if err != nil {
		panic(err)
	}
	return &post, nil
}


func (table *dynamoDBTable) DeleteById(id string) (*models.Post, error) {
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	_, err := dynamoDBClient.DeleteItem(c, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *aws.String(id)},
		},
	})
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

