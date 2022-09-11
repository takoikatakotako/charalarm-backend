package repository

import (
	"charalarm/entity"
	"fmt"
	"context"
	// "github.com/aws/aws-lambda-go/lambda"
	// DynamoDBつかうとこから　
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    // "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
    awsRegion string = "ap-northeast-1"
)

type DynamoDBRepository struct{
}

func (self DynamoDBRepository) createDynamoDBClient() (*dynamodb.Client, error) {
	var ctx = context.Background()

    // DynamoDB クライアントの生成
    c, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
    if err != nil {
        fmt.Printf("load aws config: %s\n", err.Error())
        return nil, err
    }

	awsEndpoint := "http://localhost:4566"

	c.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
		   URL:           awsEndpoint,
		   SigningRegion: awsRegion,
		}, nil
	 })
	 if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
		return nil, err
	 }

	 return dynamodb.NewFromConfig(c), nil
}

func (self DynamoDBRepository) InsertAnonymousUser(anonymousUser entity.AnonymousUser) {
    var err error
    var ctx = context.Background()

    client, err := self.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
	 }

    // 新規レコードの追加
    av, err := attributevalue.MarshalMap(anonymousUser)
    if err != nil {
        fmt.Printf("dynamodb marshal: %s\n", err.Error())
        return
    }
    _, err = client.PutItem(ctx, &dynamodb.PutItemInput{
        TableName: aws.String("user-table"),
        Item:      av,
    })
    if err != nil {
        fmt.Printf("put item: %s\n", err.Error())
        return
    }
}

// ここの関数動かすところから
func (self DynamoDBRepository) GetAnonymousUser(userId string) (entity.AnonymousUser, error) {
    var err error
    var ctx = context.Background()

    client, err := self.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
	 }

    // 既存レコードの取得
    getInput := &dynamodb.GetItemInput{
        TableName: aws.String("user-table"),
        Key: map[string]types.AttributeValue{
            "userId": &types.AttributeValueMemberS{
                Value: userId,
            },
        },
    }

	// 取得
    output, err := client.GetItem(ctx, getInput)
    if err != nil {
        fmt.Printf("get item: %s\n", err.Error())
        return entity.AnonymousUser{}, err
    }
    gotUser := entity.AnonymousUser{}
    err = attributevalue.UnmarshalMap(output.Item, &gotUser)
    if err != nil {
        fmt.Printf("dynamodb unmarshal: %s\n", err.Error())
        return entity.AnonymousUser{}, err
    }
    fmt.Println(gotUser)
	return gotUser, nil
}

