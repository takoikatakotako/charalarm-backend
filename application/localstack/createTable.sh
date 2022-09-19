#!/bin/bash
set -x

awslocal dynamodb create-table \
    --table-name user-table \
    --attribute-definitions AttributeName=userID,AttributeType=S \
    --key-schema AttributeName=userID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

awslocal dynamodb create-table \
    --table-name alarm-table \
    --attribute-definitions AttributeName=alarmID,AttributeType=S \
    --key-schema AttributeName=alarmID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

set +x
