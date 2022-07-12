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


var tableName = "posts"

type dynamoDBTable struct {
	tableName string
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
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: getEnv("Endpoint"), SigningRegion: getEnv("Region")}, nil
			})),
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
	attributeValue, err := attributevalue.MarshalMap(post)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(c, item)
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

// Delete: TODO
func (table *dynamoDBTable) DeleteById(id string) (*models.Post, error) {
	c := context.Background()
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.DeleteItem(c, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *aws.String(id)},
		},
	})
	if err != nil {
		return nil, err
	}
	post := models.Post{}
	err = attributevalue.UnmarshalMap(result.Attributes, &post)
	if err != nil {
		panic(err)
	}
	return &post, nil
}

