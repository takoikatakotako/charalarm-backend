package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	// "github.com/takoikatakotako/charalarm-backend/table"
	// "github.com/takoikatakotako/charalarm-backend/validator"
	charalarm_config "github.com/takoikatakotako/charalarm-backend/config"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

type SNSRepository struct {
	IsLocal bool
}

func (s *SNSRepository) createSNSClient() (*sns.Client, error) {
	ctx := context.Background()

	// SNS クライアントの生成
	c, err := config.LoadDefaultConfig(ctx, config.WithRegion(charalarm_config.AWSRegion))
	if err != nil {
		fmt.Printf("load aws config: %s\n", err.Error())
		return nil, err
	}

	// LocalStackを使う場合
	if s.IsLocal {
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
	return sns.NewFromConfig(c), nil
}

////////////////////////////////////
// iOS Platform Endpoint
////////////////////////////////////
func (s *SNSRepository) CreateIOSPushPlatformEndpoint(pushToken string) (entity.CreatePlatformEndpointResponse, error) {
	platformApplicationArn := "arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application"
	return s.createPlatformEndpoint(platformApplicationArn, pushToken)
}

func (s *SNSRepository) CreateIOSVoipPushPlatformEndpoint(pushToken string) (entity.CreatePlatformEndpointResponse, error) {
	platformApplicationArn := "arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application"
	return s.createPlatformEndpoint(platformApplicationArn, pushToken)
}

func (s *SNSRepository) createPlatformEndpoint(platformApplicationArn string, pushToken string) (entity.CreatePlatformEndpointResponse, error) {
	client, err := s.createSNSClient()
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}

	// エンドポイント作成
	getInput := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(platformApplicationArn),
		Token:                  aws.String(pushToken),
	}
	result, err := client.CreatePlatformEndpoint(context.Background(), getInput)
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}

	response := entity.CreatePlatformEndpointResponse{EndpointArn: *result.EndpointArn}
	return response, nil
}

func (s *SNSRepository) PublishPlatformApplication(alarmInfo entity.AlarmInfo) error {
	client, err := s.createSNSClient()
	if err != nil {
		return err
	}

	// エンドポイントを取得
	getEndpointAttributesInput := &sns.GetEndpointAttributesInput{
		EndpointArn: aws.String(alarmInfo.SNSEndpointArn),
	}
	getEndpointAttributesOutput, err := client.GetEndpointAttributes(context.Background(), getEndpointAttributesInput)
	if err != nil {
		return err
	}

	isEnabled := getEndpointAttributesOutput.Attributes["Enabled"]
	if isEnabled == "False" || isEnabled == "false" {
		return errors.New("EndpointがFalse")
	}

	// メッセージを作成
	voipPushInfo := entity.VoIPPushInfo{}
	voipPushInfo.CharaName = alarmInfo.CharaName
	voipPushInfo.FilePath = alarmInfo.FileURL
	jsonBytes, err := json.Marshal(voipPushInfo)
	if err != nil {
		return err
	}

	// プッシュ通知
	publishInput := &sns.PublishInput{
		Message:   aws.String(string(jsonBytes)),
		TargetArn: aws.String(alarmInfo.SNSEndpointArn),
	}
	_, err = client.Publish(context.Background(), publishInput)
	if err != nil {
		return err
	}

	return nil
}
