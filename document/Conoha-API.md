## Conoha で Charalarm API を構築

### Pull Image

```
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com
docker pull 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-api:latest
```

### Docker でアプリを起動

`/opt/charalarm-api` に `docker-compose.yml` ファイルを作成する。 `AWS_SNS_URL` は local-stack 用のものなのでサーバーで動かすときはダミーの値を入れる

```
version: "3.8"
services:
  charalarm-api:
    image: 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-api:latest
    restart: always
    ports:
      - 8080:8080
    environment:
      SPRING_DATASOURCE_URL: jdbc:mysql://prd-charalarm-mysql-database.cgm0celpahsu.ap-northeast-1.rds.amazonaws.com/stg_charalarm
      SPRING_DATASOURCE_USERNAME: stg_charalarm
      SPRING_DATASOURCE_PASSWORD: STG_MYSQL_PASSWORD
      AWS_SNS_URL: "DAMMY_VALUE"
      AWS_SNS_IOS_VOIP_PUSH_ARN: arn:aws:sns:ap-northeast-1:071766756112:app/APNS_VOIP/prd-charalarm-ios-voip-push
      AWS_SNS_IOS_PUSH_ARN: arn:aws:sns:ap-northeast-1:071766756112:app/APNS/prd-charalarm-ios-push
```
