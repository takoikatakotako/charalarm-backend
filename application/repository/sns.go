package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"strings"

	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	// "github.com/takoikatakotako/charalarm-backend/table"
	// "github.com/takoikatakotako/charalarm-backend/validator"
	charalarm_config "github.com/takoikatakotako/charalarm-backend/config"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

const (
	iOSPushPlatformApplication     = "ios-push-platform-application"
	iOSVoIPPushPlatformApplication = "ios-voip-push-platform-application"
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

// CreateIOSPushPlatformEndpoint iOS Platform Endpoint
func (s *SNSRepository) CreateIOSPushPlatformEndpoint(pushToken string) (entity.CreatePlatformEndpointResponse, error) {
	platformApplicationArn, err := s.getPlatformApplicationARN(iOSPushPlatformApplication)
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}
	return s.createPlatformEndpoint(platformApplicationArn, pushToken)
}

func (s *SNSRepository) CreateIOSVoipPushPlatformEndpoint(pushToken string) (entity.CreatePlatformEndpointResponse, error) {
	platformApplicationArn, err := s.getPlatformApplicationARN(iOSVoIPPushPlatformApplication)
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}
	return s.createPlatformEndpoint(platformApplicationArn, pushToken)
}

func (s *SNSRepository) CheckPlatformEndpointEnabled(endpoint string) error {
	client, err := s.createSNSClient()
	if err != nil {
		return err
	}

	// エンドポイントを取得
	getEndpointAttributesInput := &sns.GetEndpointAttributesInput{
		EndpointArn: aws.String(endpoint),
	}
	getEndpointAttributesOutput, err := client.GetEndpointAttributes(context.Background(), getEndpointAttributesInput)
	if err != nil {
		return err
	}

	isEnabled := getEndpointAttributesOutput.Attributes["Enabled"]
	if isEnabled == "False" || isEnabled == "false" {
		return errors.New("EndpointがFalse")
	}

	return nil
}

func (s *SNSRepository) PublishPlatformApplication(targetArn string, message string) error {
	client, err := s.createSNSClient()
	if err != nil {
		return err
	}

	// プッシュ通知を発火
	publishInput := &sns.PublishInput{
		Message:   aws.String(message),
		TargetArn: aws.String(targetArn),
	}
	_, err = client.Publish(context.Background(), publishInput)
	if err != nil {
		return err
	}

	return nil
}

//////////////////////////////
// Private Methods
//////////////////////////////

// getQueueURL QueueのURLを取得する
func (s *SNSRepository) getPlatformApplicationARN(queueName string) (string, error) {
	// SQSClient作成
	client, err := s.createSNSClient()
	if err != nil {
		return "", err
	}

	// PlatformApplication を取得
	input := &sns.ListPlatformApplicationsInput{}
	output, err := client.ListPlatformApplications(context.Background(), input)
	if err != nil {
		return "", err
	}

	for _, platformApplication := range output.PlatformApplications {
		platformApplicationArn := *platformApplication.PlatformApplicationArn
		if strings.Contains(platformApplicationArn, queueName) {
			return platformApplicationArn, nil
		}
	}

	return "", errors.New("platform Application Not Found")
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
