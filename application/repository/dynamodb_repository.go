package repository

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/table"
	"context"
	"errors"
	"fmt"
    "encoding/json"
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
func (self DynamoDBRepository) GetAnonymousUser(userID string) (entity.AnonymousUser, error) {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return entity.AnonymousUser{}, err
	}

	// 既存レコードの取得
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(table.USER_TABLE),
		Key: map[string]types.AttributeValue{
			"userID": &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	}

	// 取得
	output, err := client.GetItem(ctx, getInput)
	if err != nil {
		return entity.AnonymousUser{}, err
	}
	gotUser := entity.AnonymousUser{}

	if len(output.Item) == 0 {
		return entity.AnonymousUser{}, errors.New(charalarm_error.INVAlID_VALUE)
	}

	err = attributevalue.UnmarshalMap(output.Item, &gotUser)
	if err != nil {
		return entity.AnonymousUser{}, err
	}

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
		TableName: aws.String(table.USER_TABLE),
		Key: map[string]types.AttributeValue{
			"userID": &types.AttributeValueMemberS{
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

func (self DynamoDBRepository) InsertAnonymousUser(anonymousUser entity.AnonymousUser) error {
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
		TableName: aws.String(table.USER_TABLE),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

func (self DynamoDBRepository) DeleteAnonymousUser(userID string) error {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return err
	}

	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(table.USER_TABLE),
		Key: map[string]types.AttributeValue{
			"userID": &types.AttributeValueMemberS{
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

////////////////////////////////////
// Alarm
////////////////////////////////////
func (self DynamoDBRepository) GetAlarmList(userID string) ([]entity.Alarm, error) {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return []entity.Alarm{}, err
	}

	// クエリ実行
	output, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("alarm-table"),
		IndexName:              aws.String("user-id-index"),
		KeyConditionExpression: aws.String("userID = :userID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userID": &types.AttributeValueMemberS{Value: userID},
		},
	})
	if err != nil {
		return []entity.Alarm{}, err
	}

	// 取得結果を struct の配列に変換
	alarmList := []entity.Alarm{}
	for _, item := range output.Items {
		alarm := entity.Alarm{}
		err = attributevalue.UnmarshalMap(item, &alarm)
		if err != nil {
			return []entity.Alarm{}, err
		}
		alarmList = append(alarmList, alarm)
	}

	return alarmList, nil
}

func (self DynamoDBRepository) InsertAlarm(alarm entity.Alarm) error {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// Alarm のバリデーション

	// 新規レコードの追加
	av, err := attributevalue.MarshalMap(alarm)
	if err != nil {
		fmt.Printf("dynamodb marshal: %s\n", err.Error())
		return err
	}
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.ALARM_TABLE),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

func (self DynamoDBRepository) DeleteAlarm(alarmID string) error {
	var err error
	var ctx = context.Background()

	client, err := self.createDynamoDBClient()
	if err != nil {
		return err
	}

    xx, err := client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
        TableName: aws.String(table.ALARM_TABLE),
        Key: map[string]types.AttributeValue{
            "alarmID": &types.AttributeValueMemberS{Value: alarmID},
        },
    })
	if err != nil {
		return err
	}

	fmt.Println("------")
	fmt.Println(alarmID)
    bs, _ := json.Marshal(xx)
    fmt.Println(string(bs))
	fmt.Println("------")

	return nil
}
