package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	// // "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/takoikatakotako/charalarm-backend/entity"
	// // charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
	// // "github.com/takoikatakotako/charalarm-backend/table"
	// // "github.com/takoikatakotako/charalarm-backend/validator"
)

type SQSRepository struct {
	IsLocal bool
}

func (s *SQSRepository) createSQSClient() (*sqs.Client, error) {
	ctx := context.Background()

	// SQS クライアントの生成
	c, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		fmt.Printf("load aws config: %s\n", err.Error())
		return nil, err
	}

	// LocalStackを使う場合
	if s.IsLocal {
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
	return sqs.NewFromConfig(c), nil
}

////////////////////////////////////
// SQS
////////////////////////////////////
func (s *SQSRepository) SendAlarmInfoMessage(alarmInfo entity.AlarmInfo) (error) {
	queueURL := "http://localhost:4566/000000000000/voip-push-queue.fifo"
	return s.sendMessage(queueURL, alarmInfo)
}

func (s *SQSRepository) sendMessage(queueURL string, alarmInfo entity.AlarmInfo) (error) {

	// SQSClient作成
	client, err := s.createSQSClient()
	if err != nil {
		return err
	}

	sMInput := &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{},
		MessageGroupId: aws.String("XXXX"),
		MessageDeduplicationId: aws.String("XXXX"),
		MessageBody: aws.String("In bestseller for the week of 12/11/2016."),
		QueueUrl:    aws.String(queueURL),
	}


	resp, err := client.SendMessage(context.TODO(), sMInput)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return err
	}

	fmt.Println("Sent message with ID: " + *resp.MessageId)


	return nil
}

// func (s *SNSRepository) CreateIOSVoipPushPlatformEndpoint(pushToken string) (entity.CreatePlatformEndpointResponse, error) {
// 	platformApplicationArn := "arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application"
// 	return s.createPlatformEndpoint(platformApplicationArn, pushToken)
// }

// func (s *SNSRepository) createPlatformEndpoint(platformApplicationArn string, pushToken string) (entity.CreatePlatformEndpointResponse, error) {
// 	ctx := context.Background()

// 	client, err := s.createSNSClient()
// 	if err != nil {
// 		return entity.CreatePlatformEndpointResponse{}, err
// 	}

// 	// エンドポイント作成
// 	getInput := &sns.CreatePlatformEndpointInput{
// 		PlatformApplicationArn: aws.String(platformApplicationArn),
// 		Token:                  aws.String(pushToken),
// 	}
// 	result, err := client.CreatePlatformEndpoint(ctx, getInput)
// 	if err != nil {
// 		return entity.CreatePlatformEndpointResponse{}, err
// 	}

// 	response := entity.CreatePlatformEndpointResponse{EndpointArn: *result.EndpointArn}
// 	return response, nil
// }

// // func (d *DynamoDBRepository) IsExistAnonymousUser(userID string) (bool, error) {
// // 	ctx := context.Background()

// // 	// DBClient作成
// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		return false, err
// // 	}

// // 	// 既存レコードの取得
// // 	getInput := &dynamodb.GetItemInput{
// // 		TableName: aws.String(table.USER_TABLE),
// // 		Key: map[string]types.AttributeValue{
// // 			"userID": &types.AttributeValueMemberS{
// // 				Value: userID,
// // 			},
// // 		},
// // 	}
// // 	response, err := client.GetItem(ctx, getInput)
// // 	if err != nil {
// // 		return false, err
// // 	}

// // 	if len(response.Item) == 0 {
// // 		return false, nil
// // 	} else {
// // 		return true, nil
// // 	}
// // }

// // func (d *DynamoDBRepository) InsertAnonymousUser(anonymousUser entity.AnonymousUser) error {
// // 	ctx := context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		fmt.Printf("err, %v", err)
// // 		return err
// // 	}

// // 	// 新規レコードの追加
// // 	av, err := attributevalue.MarshalMap(anonymousUser)
// // 	if err != nil {
// // 		fmt.Printf("dynamodb marshal: %s\n", err.Error())
// // 		return err
// // 	}
// // 	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
// // 		TableName: aws.String(table.USER_TABLE),
// // 		Item:      av,
// // 	})
// // 	if err != nil {
// // 		fmt.Printf("put item: %s\n", err.Error())
// // 		return err
// // 	}

// // 	return nil
// // }

// // func (d *DynamoDBRepository) DeleteAnonymousUser(userID string) error {
// // 	ctx := context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		return err
// // 	}

// // 	deleteInput := &dynamodb.DeleteItemInput{
// // 		TableName: aws.String(table.USER_TABLE),
// // 		Key: map[string]types.AttributeValue{
// // 			"userID": &types.AttributeValueMemberS{
// // 				Value: userID,
// // 			},
// // 		},
// // 	}

// // 	_, err = client.DeleteItem(ctx, deleteInput)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // ////////////////////////////////////
// // // Alarm
// // ////////////////////////////////////
// // func (d *DynamoDBRepository) GetAlarmList(userID string) ([]entity.Alarm, error) {
// // 	ctx := context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		return []entity.Alarm{}, err
// // 	}

// // 	// クエリ実行
// // 	output, err := client.Query(ctx, &dynamodb.QueryInput{
// // 		TableName:              aws.String("alarm-table"),
// // 		IndexName:              aws.String("user-id-index"),
// // 		KeyConditionExpression: aws.String("userID = :userID"),
// // 		ExpressionAttributeValues: map[string]types.AttributeValue{
// // 			":userID": &types.AttributeValueMemberS{Value: userID},
// // 		},
// // 	})
// // 	if err != nil {
// // 		return []entity.Alarm{}, err
// // 	}

// // 	// 取得結果を struct の配列に変換
// // 	alarmList := []entity.Alarm{}
// // 	for _, item := range output.Items {
// // 		alarm := entity.Alarm{}
// // 		if err := attributevalue.UnmarshalMap(item, &alarm); err != nil {
// // 			return []entity.Alarm{}, err
// // 		}
// // 		alarmList = append(alarmList, alarm)
// // 	}

// // 	return alarmList, nil
// // }

// // func (d *DynamoDBRepository) InsertAlarm(alarm entity.Alarm) error {
// // 	ctx := context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		fmt.Printf("err, %v", err)
// // 		return err
// // 	}

// // 	// Alarm のバリデーション
// // 	if !validator.IsValidateAlarm(alarm) {
// // 		return errors.New(charalarm_error.INVAlID_VALUE)
// // 	}

// // 	// 新規レコードの追加
// // 	av, err := attributevalue.MarshalMap(alarm)
// // 	if err != nil {
// // 		fmt.Printf("dynamodb marshal: %s\n", err.Error())
// // 		return err
// // 	}
// // 	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
// // 		TableName: aws.String(table.ALARM_TABLE),
// // 		Item:      av,
// // 	})
// // 	if err != nil {
// // 		fmt.Printf("put item: %s\n", err.Error())
// // 		return err
// // 	}

// // 	return nil
// // }

// // func (d *DynamoDBRepository) DeleteAlarm(alarmID string) error {
// // 	ctx := context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		return err
// // 	}

// // 	_, err = client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
// // 		TableName: aws.String(table.ALARM_TABLE),
// // 		Key: map[string]types.AttributeValue{
// // 			"alarmID": &types.AttributeValueMemberS{Value: alarmID},
// // 		},
// // 	})
// // 	if err != nil {
// // 		return err
// // 	}

// // 	// fmt.Println("------")
// // 	// fmt.Println(alarmID)
// // 	// bs, _ := json.Marshal(xx)
// // 	// fmt.Println(string(bs))
// // 	// fmt.Println("------")

// // 	return nil
// // }

// // func (d *DynamoDBRepository) DeleteUserAlarm(userID string) error {
// // 	var err error
// // 	var ctx = context.Background()

// // 	client, err := d.createDynamoDBClient()
// // 	if err != nil {
// // 		return err
// // 	}

// // 	// userIDからアラームを検索
// // 	output, err := client.Query(ctx, &dynamodb.QueryInput{
// // 		TableName:              aws.String(table.ALARM_TABLE),
// // 		IndexName:              aws.String(table.USER_ID_INDEX),
// // 		KeyConditionExpression: aws.String("userID = :userID"),
// // 		ExpressionAttributeValues: map[string]types.AttributeValue{
// // 			":userID": &types.AttributeValueMemberS{Value: userID},
// // 		},
// // 	})
// // 	if err != nil {
// // 		return err
// // 	}

// // 	// 検索結果から一括削除のためのrequestItemsを作成
// // 	requestItems := []types.WriteRequest{}
// // 	for _, item := range output.Items {
// // 		// alarmIDを取得
// // 		alarm := entity.Alarm{}
// // 		if err := attributevalue.UnmarshalMap(item, &alarm); err != nil {
// // 			return err
// // 		}
// // 		alarmID := alarm.AlarmID

// // 		// requestItemsを作成
// // 		requestItem := types.WriteRequest{
// // 			DeleteRequest: &types.DeleteRequest{
// // 				Key: map[string]types.AttributeValue{
// // 					"alarmID": &types.AttributeValueMemberS{Value: alarmID},
// // 				},
// // 			},
// // 		}
// // 		requestItems = append(requestItems, requestItem)
// // 	}

// // 	// アラームを削除
// // 	_, err = client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
// // 		RequestItems: map[string][]types.WriteRequest{
// // 			table.ALARM_TABLE: requestItems,
// // 		},
// // 	})
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }
