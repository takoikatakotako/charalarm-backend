#!/bin/bash
set -x

# user-talbe
awslocal dynamodb create-table \
    --table-name user-table \
    --attribute-definitions AttributeName=userID,AttributeType=S \
    --key-schema AttributeName=userID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1

# alarm-table
awslocal dynamodb create-table \
    --table-name alarm-table \
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

# chara-talbe
awslocal dynamodb create-table \
    --table-name chara-table \
    --attribute-definitions AttributeName=charaID,AttributeType=S \
    --key-schema AttributeName=charaID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1


## Add Chara
awslocal dynamodb put-item \
    --table-name chara-table \
    --item '{ "charaID": { "S": "com.charalarm.yui" }, "charaEnable": { "BOOL": true }, "charaName": { "S": "井上結衣" }, "charaDescription": { "S": "井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！" }, "charaProfiles": { "L": [{"M": {"title": {"S": "イラストレーター"}, "name": {"S": "さいもん"}, "url": {"S": "https://twitter.com/simon_ns"} }}, {"M": {"title": {"S": "声優"}, "name": {"S": "Mai"}, "url": {"S": "https://twitter.com/mai_mizuiro"} }} ] } }'

awslocal dynamodb put-item \
    --table-name chara-table \
    --item '{ "charaID": { "S": "com.senpu-ki-soft.momiji" }, "charaEnable": { "BOOL": true }, "charaName": { "S": "井上結衣" }, "charaDescription": { "S": "井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！" }, "charaProfiles": { "L": [{"M": {"title": {"S": "イラストレーター"}, "name": {"S": "さいもん"}, "url": {"S": "https://twitter.com/simon_ns"} }}, {"M": {"title": {"S": "声優"}, "name": {"S": "Mai"}, "url": {"S": "https://twitter.com/mai_mizuiro"} }}, {"M": {"title": {"S": "スクリプト"}, "name": {"S": "小旗ふたる！"}, "url": {"S": "https://twitter.com/Kass_kobataku"} }} ] } }'
