package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	charalarm_config "github.com/takoikatakotako/charalarm-backend/config"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/validator"
	"math/rand"
	"time"
)

type DynamoDBRepository struct {
	IsLocal bool
}

func (d *DynamoDBRepository) createDynamoDBClient() (*dynamodb.Client, error) {
	ctx := context.Background()

	// DynamoDB クライアントの生成
	c, err := config.LoadDefaultConfig(ctx, config.WithRegion(charalarm_config.AWSRegion))
	if err != nil {
		fmt.Printf("load aws config: %s\n", err.Error())
		return nil, err
	}

	// LocalStackを使う場合
	if d.IsLocal {
		c.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           charalarm_config.LocalstackEndpoint,
				SigningRegion: charalarm_config.AWSRegion,
			}, nil
		})
		if err != nil {
			fmt.Printf("unable to load SDK config, %v", err)
			return nil, err
		}
	}
	return dynamodb.NewFromConfig(c), nil
}

// GetUser Userを取得する
func (d *DynamoDBRepository) GetUser(userID string) (database.User, error) {
	ctx := context.Background()

	client, err := d.createDynamoDBClient()
	if err != nil {
		return database.User{}, err
	}

	// 既存レコードの取得
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(database.UserTableName),
		Key: map[string]types.AttributeValue{
			database.UserTableUserId: &types.AttributeValueMemberS{
				Value: userID,
			},
		},
	}

	// 取得
	output, err := client.GetItem(ctx, getInput)
	if err != nil {
		return database.User{}, err
	}
	getUser := database.User{}

	if len(output.Item) == 0 {
		return database.User{}, errors.New(message.InvalidValue)
	}

	err = attributevalue.UnmarshalMap(output.Item, &getUser)
	if err != nil {
		return database.User{}, err
	}

	return getUser, nil
}

func (d *DynamoDBRepository) IsExistAnonymousUser(userID string) (bool, error) {
	ctx := context.Background()

	// DBClient作成
	client, err := d.createDynamoDBClient()
	if err != nil {
		return false, err
	}

	// 既存レコードの取得
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(database.UserTableName),
		Key: map[string]types.AttributeValue{
			database.UserTableUserId: &types.AttributeValueMemberS{
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

func (d *DynamoDBRepository) InsertUser(anonymousUser database.User) error {
	ctx := context.Background()

	client, err := d.createDynamoDBClient()
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
		TableName: aws.String(database.UserTableName),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

func (d *DynamoDBRepository) DeleteAnonymousUser(userID string) error {
	ctx := context.Background()

	client, err := d.createDynamoDBClient()
	if err != nil {
		return err
	}

	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(database.UserTableName),
		Key: map[string]types.AttributeValue{
			database.UserTableUserId: &types.AttributeValueMemberS{
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
func (d *DynamoDBRepository) GetAlarmList(userID string) ([]database.Alarm, error) {

	client, err := d.createDynamoDBClient()
	if err != nil {
		return []database.Alarm{}, err
	}

	// クエリ実行
	keyEx := expression.Key(database.ALARM_TABLE_USER_ID).Equal(expression.Value(userID))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	output, err := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(database.AlarmTableName),
		IndexName:                 aws.String(database.ALARM_TABLE_USER_ID_INDEX_NAME),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return []database.Alarm{}, err
	}

	// 取得結果を struct の配列に変換
	alarmList := []database.Alarm{}
	for _, item := range output.Items {
		alarm := database.Alarm{}
		err := attributevalue.UnmarshalMap(item, &alarm)
		if err != nil {
			// TODO ログを出す
			continue
		}
		alarmList = append(alarmList, alarm)
	}

	return alarmList, nil
}

func (d *DynamoDBRepository) QueryByAlarmTime(hour int, minute int, weekday time.Weekday) ([]database.Alarm, error) {
	alarmTime := fmt.Sprintf("%02d-%02d", hour, minute)

	// clientの作成
	client, err := d.createDynamoDBClient()
	if err != nil {
		return []database.Alarm{}, err
	}

	keyEx := expression.Key(database.ALARM_TABLE_TIME).Equal(expression.Value(alarmTime))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()

	output, err := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(database.AlarmTableName),
		IndexName:                 aws.String(database.ALARM_TABLE_ALARM_TIME_INDEX_NAME),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return []database.Alarm{}, err
	}

	// 取得結果を struct の配列に変換
	alarmList := []database.Alarm{}
	for _, item := range output.Items {
		alarm := database.Alarm{}
		err := attributevalue.UnmarshalMap(item, &alarm)
		if err != nil {
			// TODO ログを出す
			continue
		}

		// 曜日が一致するもの
		switch weekday {
		case time.Sunday:
			if alarm.Sunday {
				alarmList = append(alarmList, alarm)
			}
		case time.Monday:
			if alarm.Monday {
				alarmList = append(alarmList, alarm)
			}
		case time.Tuesday:
			if alarm.Tuesday {
				alarmList = append(alarmList, alarm)
			}
		case time.Wednesday:
			if alarm.Wednesday {
				alarmList = append(alarmList, alarm)
			}
		case time.Thursday:
			if alarm.Thursday {
				alarmList = append(alarmList, alarm)
			}
		case time.Friday:
			if alarm.Friday {
				alarmList = append(alarmList, alarm)
			}
		case time.Saturday:
			if alarm.Saturday {
				alarmList = append(alarmList, alarm)
			}
		}
	}

	return alarmList, nil
}

func (d *DynamoDBRepository) InsertAlarm(alarm database.Alarm) error {
	client, err := d.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// Alarm のバリデーション
	err = validator.ValidateAlarm(alarm)
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// 新規レコードの追加
	av, err := attributevalue.MarshalMap(alarm)
	if err != nil {
		fmt.Printf("dynamodb marshal: %s\n", err.Error())
		return err
	}
	_, err = client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(database.AlarmTableName),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

// TODO: 少し危険な方法で更新しているので、更新対象の数だけメソッドを作成する
func (d *DynamoDBRepository) UpdateAlarm(alarm database.Alarm) error {
	client, err := d.createDynamoDBClient()
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// Alarm のバリデーション
	err = validator.ValidateAlarm(alarm)
	if err != nil {
		fmt.Printf("err, %v", err)
		return err
	}

	// レコードを更新する
	av, err := attributevalue.MarshalMap(alarm)
	if err != nil {
		fmt.Printf("dynamodb marshal: %s\n", err.Error())
		return err
	}
	_, err = client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(database.AlarmTableName),
		Item:      av,
	})
	if err != nil {
		fmt.Printf("put item: %s\n", err.Error())
		return err
	}

	return nil
}

func (d *DynamoDBRepository) DeleteAlarm(alarmID string) error {
	ctx := context.Background()

	client, err := d.createDynamoDBClient()
	if err != nil {
		return err
	}

	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(database.AlarmTableName),
		Key: map[string]types.AttributeValue{
			database.ALARM_TABLE_ALARM_ID: &types.AttributeValueMemberS{Value: alarmID},
		},
	}

	_, err = client.DeleteItem(ctx, deleteInput)
	if err != nil {
		return err
	}

	// fmt.Println("------")
	// fmt.Println(alarmID)
	// bs, _ := json.Marshal(xx)
	// fmt.Println(string(bs))
	// fmt.Println("------")

	return nil
}

func (d *DynamoDBRepository) DeleteUserAlarm(userID string) error {
	var err error
	var ctx = context.Background()

	client, err := d.createDynamoDBClient()
	if err != nil {
		return err
	}

	// userIDからアラームを検索
	output, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(database.AlarmTableName),
		IndexName:              aws.String(database.UserTableUserIdIndexName),
		KeyConditionExpression: aws.String("userID = :userID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userID": &types.AttributeValueMemberS{Value: userID},
		},
	})
	if err != nil {
		return err
	}

	// 検索結果から一括削除のためのrequestItemsを作成
	requestItems := []types.WriteRequest{}
	for _, item := range output.Items {
		// alarmIDを取得
		alarm := database.Alarm{}
		if err := attributevalue.UnmarshalMap(item, &alarm); err != nil {
			return err
		}
		alarmID := alarm.AlarmID

		// requestItemsを作成
		requestItem := types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					database.ALARM_TABLE_ALARM_ID: &types.AttributeValueMemberS{Value: alarmID},
				},
			},
		}
		requestItems = append(requestItems, requestItem)
	}

	// アラームを削除
	_, err = client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			database.AlarmTableName: requestItems,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// GetChara キャラを取得する
func (d *DynamoDBRepository) GetChara(charaID string) (database.Chara, error) {
	// クライアント作成
	client, err := d.createDynamoDBClient()
	if err != nil {
		return database.Chara{}, err
	}

	// クエリ実行
	input := &dynamodb.GetItemInput{
		TableName: aws.String(database.CharaTableName),
		Key: map[string]types.AttributeValue{
			database.CharaTableCharaID: &types.AttributeValueMemberS{
				Value: charaID,
			},
		},
	}
	resp, err := client.GetItem(context.Background(), input)
	if err != nil {
		return database.Chara{}, err
	}

	if len(resp.Item) == 0 {
		return database.Chara{}, fmt.Errorf(message.ItemNotFound)
	}

	// 取得結果をcharaに変換
	chara := database.Chara{}
	err = attributevalue.UnmarshalMap(resp.Item, &chara)
	if err != nil {
		return chara, err
	}

	return chara, nil
}

// GetCharaList キャラ一覧を取得
func (d *DynamoDBRepository) GetCharaList() ([]database.Chara, error) {
	// クライアント作成
	client, err := d.createDynamoDBClient()
	if err != nil {
		return []database.Chara{}, err
	}

	// クエリ実行
	input := &dynamodb.ScanInput{
		TableName: aws.String("chara-table"),
	}
	output, err := client.Scan(context.Background(), input)
	if err != nil {
		return []database.Chara{}, err
	}

	// 取得結果を struct の配列に変換
	charaList := []database.Chara{}
	for _, item := range output.Items {
		chara := database.Chara{}
		err := attributevalue.UnmarshalMap(item, &chara)
		if err != nil {
			// TODO ログを出す
			fmt.Println(err)
			continue
		}
		charaList = append(charaList, chara)
	}

	return charaList, nil
}

// GetRandomChara
// ランダムにキャラを1つ取得する, キャラ数が増えてきた場合は改良する
func (d *DynamoDBRepository) GetRandomChara() (database.Chara, error) {
	// クライアント作成
	client, err := d.createDynamoDBClient()
	if err != nil {
		return database.Chara{}, err
	}

	// クエリ実行
	input := &dynamodb.ScanInput{
		TableName: aws.String("chara-table"),
		Limit:     aws.Int32(5),
	}
	output, err := client.Scan(context.Background(), input)
	if err != nil {
		return database.Chara{}, err
	}

	// ランダムに1件取得
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(output.Items))
	item := output.Items[index]

	// 取得結果をcharaに変換
	chara := database.Chara{}
	err = attributevalue.UnmarshalMap(item, &chara)
	if err != nil {
		return chara, err
	}

	return chara, nil
}
