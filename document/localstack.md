# LocalStack

LocalStackのデバッグに良く使うコマンドです。

## LocalStackを作り直す

```
docker-compose down
docker-compose up -d
```


## テーブルの一覧を表示

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


## スキャン

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




