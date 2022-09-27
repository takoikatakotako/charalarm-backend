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



## DynamoDB

### テーブルの一覧を表示

```
$ aws dynamodb list-tables --endpoint-url=http://localhost:4566
```

### テーブルの作成

```
$ aws dynamodb create-table \
    --table-name user-table3 \
    --attribute-definitions AttributeName=userID,AttributeType=S \
    --key-schema AttributeName=userID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url=http://localhost:4566 | jq
```

### テーブルの詳細を表示

```
$ aws dynamodb describe-table \
    --table-name user-table \
    --endpoint-url=http://localhost:4566 | jq
```


### Itemを取得

```
$ aws dynamodb get-item \
    --table-name alarm-table \
    --key '{"alarmID": {"S": "fd5fda81-194a-488e-80f1-52b02b0d6cc9"}}' \
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


### テーブルの削除

```
$ aws dynamodb delete-table \
    --table-name user-table \ 
    --endpoint-url=http://localhost:4566
```


## SNS



```
aws sns create-platform-application \
  --name ios-push-platform-application \
  --platform APNS \
  --attributes PlatformCredential=DAMMY \
  --endpoint-url http://localhost:4575
```

```
aws sns create-platform-endpoint \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-push-platform-application \
  --token MY_TOKEN \
  --endpoint-url http://localhost:4575
```

```
aws sns list-endpoints-by-platform-application \

```

```
aws sns   delete-platform-application \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/my-topic3 \
  --endpoint-url http://localhost:4575
```

```
aws sns list-platform-applications \
  --endpoint-url http://localhost:4575
```

```
aws sns list-topics \
  --endpoint-url http://localhost:4575
```
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-push-platform-application \
  --endpoint-url http://localhost:4575

