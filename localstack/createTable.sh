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
    --item '{"charaID":{"S":"com.charalarm.yui"},"charaEnable":{"BOOL":true},"charaName":{"S":"井上結衣"},"charaDescription":{"S":"井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！"},"charaProfile":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"resources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"},{"S":"smile.png"},{"S":"comfused.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"com-charalarm-yui-0.caf"},{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-6.caf"},{"S":"com-charalarm-yui-7.caf"},{"S":"com-charalarm-yui-8.caf"},{"S":"com-charalarm-yui-9.caf"},{"S":"com-charalarm-yui-10.caf"},{"S":"com-charalarm-yui-11.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"},{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"call":{"M":{"voices":{"L":[{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"expression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"}]}}},"smile":{"M":{"images":{"L":[{"S":"smile.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"}]}}},"comfused":{"M":{"images":{"L":[{"S":"comfused.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"}]}}}}}}'

awslocal dynamodb put-item \
    --table-name chara-table \
    --item '{"charaID":{"S":"com.senpu-ki-soft.momiji"},"charaEnable":{"BOOL":true},"charaName":{"S":"紅葉"},"charaDescription":{"S":"金髪紅眼の美少女。疲れ気味のあなたを心配して様々な癒しを、と考えている。その正体は幾百年を生きる鬼の末裔。あるいはあなたに恋慕を抱く彼女。ちょっと素直になりきれないものの、なんやかんやいってそばにいてくれる面倒見のいい少女。日々あなたの生活を見届けている。「わっち？　名は紅葉でありんす。主様の支えになれるよう、掃除でもみみかきでもなんでも言っておくんなんし。か、かわいい？　い、いきなりそんなこと言わないでおくんなんし！」"},"charaProfile":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"resources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"},{"S":"tap-head-1.caf"},{"S":"tap-head-2.caf"},{"S":"tap-head-3.caf"},{"S":"tap-lower-body-1.caf"},{"S":"tap-lower-body-2.caf"},{"S":"tap-lower-body-3.caf"},{"S":"tap-upper-body-1.caf"},{"S":"tap-upper-body-2.caf"},{"S":"tap-upper-body-3.caf"},{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"call":{"M":{"voices":{"L":[{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"expression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"}]}}}}}}'
