## Conoha で Charalarm Worker を構築

#### Pull Image

```
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com
docker pull 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-worker:latest
```

### Docker でアプリを起動

```
cd opt
mkdir charalarm-worker
vi docker-compose.yml
```

`/opt/charalarm-worker` に `docker-compose.yml` ファイルを作成する。


```
version: "3.8"
services:
  charalarm-worker:
    image: 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-worker:latest
    restart: always
    environment:
      AWS_REGION=ap-northeast-1 \
      AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
      AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
      SQS_ENDPOINT: arn:aws:sqs:ap-northeast-1:071766756112:prd-voip-push-queue.fifo
      SNS_ENDPOINT: arn:aws:sns:ap-northeast-1:071766756112:app/APNS_VOIP/prd-charalarm-ios-voip-push
      SQS_QUEUE_NAME: prd-voip-push-queue.fifo
```
