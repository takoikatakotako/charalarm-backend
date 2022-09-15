package repository

import (
	"charalarm/entity"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	awsRegion string = "ap-northeast-1"
)

const (
	localstackEndpoint string = "http://localhost:4566"
)

type DynamoDBRepository struct {
	IsLocal bool
}

func (self DynamoDBRepository) createDynamoDBClient() (*dynamodb.Client, error) {
	var ctx = context.Background()

	// DynamoDB クライアントの生成
	c, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		fmt.Printf("load aws config: %s\n", err.Error())
		return nil, err
	}

	// LocalStackを使う場合
	if self.IsLocal {
		c.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           localstackEndpoint,
				SigningRegion: awsRegion,
			}, nil
		})
		if err != nil {
			fmt.Printf("unable to load SDK config, %v", err)
			return nil, err
		}
	}
	return dynamodb.NewFromConfig(c), nil
}

////////////////////////////////////
// AnonymousUser
////////////////////////////////////
func (self DynamoDBRepository) GetAnonymousUser(userId string) (entity.AnonymousUser, error) {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return entity.AnonymousUser{}, err
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
		return entity.AnonymousUser{}, err
	}
	gotUser := entity.AnonymousUser{}
	err = attributevalue.UnmarshalMap(output.Item, &gotUser)
	if err != nil {
		return entity.AnonymousUser{}, err
	}
	fmt.Println(gotUser)
	return gotUser, nil
}

func (self DynamoDBRepository) IsExistAnonymousUser(userID string) (bool, error) {
	var err error
	var ctx = context.Background()

	// DBClient作成
	client, err := self.createDynamoDBClient()
	if err != nil {
		return false, err
	}

	// 既存レコードの取得
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String("user-table"),
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	}
	response, err := client.GetItem(ctx, getInput)
	if err != nil {
		return false, err
	}

	if len(response.Item) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (self DynamoDBRepository) InsertAnonymousUser(anonymousUser entity.AnonymousUser) (error) {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// 新規レコードの追加
	av, err := attributevalue.MarshalMap(anonymousUser)
	if err != nil {
		fmt.Printf("dynamodb marshal: %s\n", err.Error())
		return err
	}
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("user-table"),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

func (self DynamoDBRepository) DeleteAnonymousUser(userID string) (error) {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return err
	}

	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String("user-table"),
		Key: map[string]types.AttributeValue{
			"userId": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	}
	
	_, err = client.DeleteItem(ctx, deleteInput)
	if err != nil {
		return err
	}

	return nil
}