package repository

import (
	"context"
	// "errors"
	"fmt"

	// "encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/takoikatakotako/charalarm-backend/entity"
	// "github.com/takoikatakotako/charalarm-backend/table"
	// "github.com/takoikatakotako/charalarm-backend/validator"
)

type SNSRepository struct {
	IsLocal bool
}

func (s *SNSRepository) createSNSClient() (*sns.Client, error) {
	ctx := context.Background()

	// SNS クライアントの生成
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
	ctx := context.Background()

	client, err := s.createSNSClient()
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}

	// エンドポイント作成
	getInput := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(platformApplicationArn),
		Token:                  aws.String(pushToken),
	}
	result, err := client.CreatePlatformEndpoint(ctx, getInput)
	if err != nil {
		return entity.CreatePlatformEndpointResponse{}, err
	}

	response := entity.CreatePlatformEndpointResponse{EndpointArn: *result.EndpointArn}
	return response, nil
}
