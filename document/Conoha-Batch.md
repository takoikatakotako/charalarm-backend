## Conoha で Charalarm Batch を構築

### Pull Image

```
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com
docker pull 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-batch:latest
```

### cron で実行するファイルを作成

```
cd opt
mkdir charalarm-batch
vi charalarm-batch.sh
```

```
#!/bin/sh

docker run --rm \
  -e DATABASE_URL=jdbc:mysql://prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com/stg_charalarm \
  -e DATABASE_USER=stg_charalarm \
  -e DATABASE_PASSWORD=${DATABASE_PASSWORD} \
  -e AWS_REGION=ap-northeast-1 \
  -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
  -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
  -e SQS_ENDPOINT=DAMMY \
  -e SQS_QUEUE_NAME=prd-voip-push-queue.fifo \
  071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-batch:latest
```

実行権限を付与

```
chmod +x /opt/charalarm-batch/charalarm-batch.sh
```

`charalarm-batch.sh` を実行する cron を作成。一番下の行に追記する。

```
crontab -e
```

```
* * * * * /opt/charalarm-batch/charalarm-batch.sh
```

設定した Cron を確認する

```
crontab -l
```

