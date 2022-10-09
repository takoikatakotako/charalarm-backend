# LocalStack

LocalStackのデバッグに良く使うコマンドです。

## LocalStackを作り直す

```
docker-compose down
docker-compose up -d
```


# S3

```
aws s3 ls --endpoint-url=http://localhost:4572
```



# DynamoDB

## テーブル一覧を表示

```
$ aws dynamodb list-tables --endpoint-url=http://localhost:4566
```

## テーブルの作成

```
$ aws dynamodb create-table \
    --table-name user-table3 \
    --attribute-definitions AttributeName=userID,AttributeType=S \
    --key-schema AttributeName=userID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url=http://localhost:4566 | jq
```

## テーブルの詳細を表示

```
$ aws dynamodb describe-table \
    --table-name user-table \
    --endpoint-url=http://localhost:4566 | jq
```


## Itemを取得

### アラームを取得

```
$ aws dynamodb get-item \
    --table-name alarm-table \
    --key '{"alarmID": {"S": "fd5fda81-194a-488e-80f1-52b02b0d6cc9"}}' \
    --endpoint-url=http://localhost:4566 | jq
```

### キャラ（com.charalarm.yui）を取得

```
$ aws dynamodb get-item \
    --table-name chara-table \
    --key '{"charaID": {"S": "com.charalarm.yui"}}' \
    --endpoint-url=http://localhost:4566 | jq
```

### キャラ（com.senpu-ki-soft.momiji）を取得

```
$ aws dynamodb get-item \
    --table-name chara-table \
    --key '{"charaID": {"S": "com.senpu-ki-soft.momiji"}}' \
    --endpoint-url=http://localhost:4566 | jq
```


## Itemを追加

### キャラ（com.charalarm.yui）を追加

```
$ aws dynamodb put-item \
    --table-name chara-table \
    --item '{"charaID":{"S":"com.charalarm.yui"},"charaEnable":{"BOOL":true},"charaName":{"S":"井上結衣"},"charaDescription":{"S":"井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！"},"charaProfile":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"resources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"},{"S":"smile.png"},{"S":"comfused.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"com-charalarm-yui-0.caf"},{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-6.caf"},{"S":"com-charalarm-yui-7.caf"},{"S":"com-charalarm-yui-8.caf"},{"S":"com-charalarm-yui-9.caf"},{"S":"com-charalarm-yui-10.caf"},{"S":"com-charalarm-yui-11.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"},{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"call":{"M":{"voices":{"L":[{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"expression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"}]}}},"smile":{"M":{"images":{"L":[{"S":"smile.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"}]}}},"comfused":{"M":{"images":{"L":[{"S":"comfused.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"}]}}}}}}' \
    --endpoint-url=http://localhost:4566 | jq
```

### キャラ（com.senpu-ki-soft.momiji）を追加

```
$ aws dynamodb put-item \
    --table-name chara-table \
    --item '{"charaID":{"S":"com.senpu-ki-soft.momiji"},"charaEnable":{"BOOL":true},"charaName":{"S":"紅葉"},"charaDescription":{"S":"金髪紅眼の美少女。疲れ気味のあなたを心配して様々な癒しを、と考えている。その正体は幾百年を生きる鬼の末裔。あるいはあなたに恋慕を抱く彼女。ちょっと素直になりきれないものの、なんやかんやいってそばにいてくれる面倒見のいい少女。日々あなたの生活を見届けている。「わっち？　名は紅葉でありんす。主様の支えになれるよう、掃除でもみみかきでもなんでも言っておくんなんし。か、かわいい？　い、いきなりそんなこと言わないでおくんなんし！」"},"charaProfile":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"resources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"},{"S":"tap-head-1.caf"},{"S":"tap-head-2.caf"},{"S":"tap-head-3.caf"},{"S":"tap-lower-body-1.caf"},{"S":"tap-lower-body-2.caf"},{"S":"tap-lower-body-3.caf"},{"S":"tap-upper-body-1.caf"},{"S":"tap-upper-body-2.caf"},{"S":"tap-upper-body-3.caf"},{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"call":{"M":{"voices":{"L":[{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"expression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"}]}}}}}}' \
    --endpoint-url=http://localhost:4566 | jq
```


### クエリ

```
$ aws dynamodb query \
    --table-name alarm-table \
    --index-name user-id-index \
    --key-condition-expression "userID = :userID" \
    --expression-attribute-values '{ ":userID": { "S": "b87e945d-8912-4276-99f7-e636d7660093" } }' \
    --endpoint-url=http://localhost:4566
```

```
$ aws dynamodb query \
    --table-name alarm-table \
    --index-name alarm-time-index \
    --key-condition-expression "alarmTime = :alarmTime" \
    --expression-attribute-values '{ ":alarmTime": { "S": "XXXXX" } }' \
    --endpoint-url=http://localhost:4566
```


### スキャン

```
$ aws dynamodb scan \
    --table-name user-table \
    --endpoint-url=http://localhost:4566 | jq
```

```
$ aws dynamodb scan \
    --table-name alarm-table \
    --endpoint-url=http://localhost:4566 | jq
```

```
$ aws dynamodb scan \
    --table-name chara-table \
    --endpoint-url=http://localhost:4566 | jq
```

### テーブルの削除

```
$ aws dynamodb delete-table \
    --table-name user-table \ 
    --endpoint-url=http://localhost:4566
```


## SNS

### PlatformApplicationを作成

```
$ aws sns create-platform-application \
    --name ios-voip-push-platform-application \
    --platform APNS \
    --attributes PlatformCredential=DAMMY \
    --endpoint-url http://localhost:4566 | jq
```

### PlatformApplicationの一覧を表示

```
$ aws sns list-platform-applications \
    --endpoint-url http://localhost:4566 | jq
```

```
{
  "PlatformApplications": [
    {
      "PlatformApplicationArn": "arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application",
      "Attributes": {
        "PlatformCredential": "DAMMY"
      }
    }
  ]
}
```


### PlatformEndpointを作成

```
aws sns create-platform-endpoint \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application \
  --token MY_TOKEN2 \
  --endpoint-url http://localhost:4566 | jq
```

### PlatformEndpointのEndpointsの一覧を確認

```
aws sns list-endpoints-by-platform-application \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-voip-push-platform-application \
  --endpoint-url http://localhost:4566

```

```
aws sns delete-platform-application \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/my-topic3 \
  --endpoint-url http://localhost:4566
```


```
aws sns list-topics \
  --endpoint-url http://localhost:4566
```
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-push-platform-application \
  --endpoint-url http://localhost:4566


# SQS

### キューの一覧を表示

```
$ aws sqs list-queues \
    --endpoint-url http://localhost:4566 | jq
```

### キューのメッセージ数を確認

```
$ aws sqs get-queue-attributes \
    --queue-url http://localhost:4566/000000000000/voip-push-queue.fifo \
    --attribute-names ApproximateNumberOfMessages \
    --endpoint-url http://localhost:4566 | jq
```

### キューのNotVisible状態のメッセージ数を確認

```
$ aws sqs get-queue-attributes \
    --queue-url http://localhost:4566/000000000000/voip-push-queue.fifo \
    --attribute-names ApproximateNumberOfMessagesNotVisible \
    --endpoint-url http://localhost:4566 | jq
```

### キューからメッセージを取得

```
$ aws sqs receive-message \
    --queue-url http://localhost:4566/000000000000/voip-push-queue.fifo \
    --max-number-of-messages 10 \
    --endpoint-url http://localhost:4566 | jq
```

### キュー内のメッセージをすべて削除

```
$ aws sqs purge-queue \
    --queue-url http://localhost:4566/000000000000/voip-push-queue.fifo \
    --endpoint-url http://localhost:4566 | jq
```
