#!/bin/bash
set -x

# user-talbe
awslocal dynamodb create-table \
    --table-name user-table \
    --attribute-definitions AttributeName=userID,AttributeType=S \
    --key-schema AttributeName=userID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

# alarm-table
aws dynamodb create-table \
    --table-name alarm-table8 \
    --attribute-definitions AttributeName=alarmID,AttributeType=S \
                            AttributeName=userID,AttributeType=S \
                            AttributeName=alarmTime,AttributeType=S  \
    --key-schema AttributeName=alarmID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --global-secondary-indexes \
        "[
            {
                \"IndexName\": \"user-id-index\",
                \"KeySchema\": [{\"AttributeName\":\"userID\",\"KeyType\":\"HASH\"}],
                \"Projection\":{
                    \"ProjectionType\":\"ALL\"
                },
                \"ProvisionedThroughput\": {
                    \"ReadCapacityUnits\": 1,
                    \"WriteCapacityUnits\": 1
                }
            },
            {
                \"IndexName\": \"alarm-time-index\",
                \"KeySchema\": [{\"AttributeName\":\"alarmTime\",\"KeyType\":\"HASH\"}],
                \"Projection\":{
                    \"ProjectionType\":\"ALL\"
                },
                \"ProvisionedThroughput\": {
                    \"ReadCapacityUnits\": 1,
                    \"WriteCapacityUnits\": 1
                }
            }
        ]"

set +x
