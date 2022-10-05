package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/takoikatakotako/charalarm-backend/entity"
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
func (s *SQSRepository) SendAlarmInfoMessage(alarmInfo entity.AlarmInfo) error {
	queueURL := "http://localhost:4566/000000000000/voip-push-queue.fifo"
	return s.sendMessage(queueURL, "XXXX", alarmInfo)
}

func (s *SQSRepository) RecieveAlarmInfoMessage() ([]types.Message, error) {
	queueURL := "http://localhost:4566/000000000000/voip-push-queue.fifo"
	return s.recieveMessage(queueURL)
}

func (s *SQSRepository) PurgeQueue() error {
	queueURL := "http://localhost:4566/000000000000/voip-push-queue.fifo"

	// SQSClient作成
	client, err := s.createSQSClient()
	if err != nil {
		return err
	}

	// purge queue
	input := &sqs.PurgeQueueInput{
		QueueUrl: aws.String(queueURL),
	}
	_, err = client.PurgeQueue(context.Background(), input)
	return err
}

// Private Methods
func (s *SQSRepository) sendMessage(queueURL string, messageGroupId string, alarmInfo entity.AlarmInfo) error {
	// SQSClient作成
	client, err := s.createSQSClient()
	if err != nil {
		return err
	}

	// decode
	jsonBytes, err := json.Marshal(alarmInfo)
	if err != nil {
		return err
	}

	// sent message
	sMInput := &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{},
		MessageGroupId:    aws.String(messageGroupId),
		MessageBody:       aws.String(string(jsonBytes)),
		QueueUrl:          aws.String(queueURL),
	}
	_, err = client.SendMessage(context.Background(), sMInput)
	return err
}

func (s *SQSRepository) recieveMessage(queueURL string) ([]types.Message, error) {
	// SQSClient作成
	client, err := s.createSQSClient()
	if err != nil {
		return []types.Message{}, nil
	}

	// receive message
	timeout := 5
	gMInput := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   int32(timeout),
	}

	resp, err := client.ReceiveMessage(context.Background(), gMInput)
	if err != nil {
		return []types.Message{}, err
	}

	return resp.Messages, nil
}
