#!/bin/bash
set -x
awslocal sqs create-queue --queue-name voip-push-queue.fifo --region ap-northeast-1



awslocal dynamodb create-table \
    --table-name user-table \
    --attribute-definitions AttributeName=Artist,AttributeType=S AttributeName=SongTitle,AttributeType=S \
    --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5


set +x
