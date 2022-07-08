package dynamoapi

//Creates the initial table on local DynamoDb
//AWS CLI: aws  --endpoint-url http://localhost:8000 dynamodb create-table --cli-input-json file://users.json

import (
	"fmt"
	"log"
	"context"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/aws"
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

  //viperenv := viperEnvVariable("STRONGEST_AVENGER")

  func DBcon() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
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

  func TestClient(){
	tableName := ""
	input := dynamodb.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	fmt.Printf("client data: %+v\n", input)
  }