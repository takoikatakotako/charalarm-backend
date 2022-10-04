# CharalarmBackend


## ドキュメント

- [APIのエンドポイントについて](document/api-endpoint.md)
- [Curlについて](document/curl.md)
- [Localstackについて](document/localstack.md)
- [Conohaサーバーのセットアップについて](documents/Conoha-Setup.md)
- [ConohaサーバーのAPIセットアップについて](documents/Conoha-API.md)
- [ConohaサーバーのBatchセットアップについて](documents/Conoha-Batch.md)
- [ConohaサーバーのWorkerセットアップについて](documents/Conoha-Worker.md)
- [Conohaサーバーのアップデートについて](documents/Conoha-Update.md)







	// fmt.Println("------")
	// bs, _ := json.Marshal(output)
	// fmt.Println(string(bs))
	// fmt.Println("------")







### API Server

Charalarm の API です。

### Alarm Batch

毎分起動するアラーム用のバッチです。送るべきアラームを取得し、加工し、SQSにプッシュします。

1. アラームテーブルから送るべきアラームを取得する。
2. アラームの送り先のSNS ARNを取得する。
3. 良い感じに加工してSQSにプッシュする

### Worker

常時稼働しているWorkerです。
SQSの情報を取得し、良い感じにパースし、APNs(Apple Push Notification Service)に電話をかけるように依頼します。


## Databaseとデータの流れ

### 匿名ユーザー(v1で実装)

1. アプリをダウンロードし、初回起動時に匿名ユーザー名とトークンがアプリ内で発行される。共にUUID。
2. 発行された匿名ユーザー名とトークンを送りサインアップを行う。
3. アラームを追加したり削除したりする。APIを叩く度に匿名ユーザー名とパスワードをリクエストに含める。(セキュリティ的には良くない気がするが、自動生成したパスワードなのと、匿名ユーザーで投稿とかはさせる気がないのでセーフの認識)

### 匿名ユーザー -> 認証済みユーザー(v2で実装、未実装)

1. Cognitoにサインアップし、AuthToken と Cognitoユーザー名を取得する。
2. SignUpAPIを叩き、認証済みユーザーを作成する。匿名ユーザー名とCognito を紐づける。紐づけるというより、匿名ユーザー時に作ったデータを新しく作った認証済みユーザーの方に移行する。

![er](./material/er.png)


## 環境構築

header に APIバージョンを追加

```
curl -H "X-API-VERSION: 0" http://localhost:8080/api/news/list
```


イメージのビルド

```
docker build -t takoikatakotako/charalarm:latest .
docker run -d -p 8080:8080 -e TWILIO_ACCOUNT_SID=xxxx -e TWILIO_AUTH_TOKEN=xxxxx takoikatakotako/charalarm:latest
```

イメージの中に入る

```
docker run --rm -it charalarm-api /bin/sh
```

テスト実行

```
./gradlew test -i
./gradlew cleanTest
```

ECR へのプッシュ

```
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com
docker build -t charalarm-api .
docker tag charalarm-api:latest 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-api:latest
docker push 071766756112.dkr.ecr.ap-northeast-1.amazonaws.com/charalarm-api:latest
```

DBを作り直す

```
docker-compose up -d

mysql --host 127.0.0.1 --port 3306 -u root -p

drop database charalarm;
create database charalarm;
use charalarm;
```

ER図を作成（TODO: Docker化したいね）

```
brew install graphviz --with-librsvg --with-pango
cd sql
java -jar schemaspy-6.1.0.jar -t mysql -dp ./schemaspy/drivers/mysql-connector-java-8.0.20.jar -host 127.0.0.1 -port 3306 -db charalarm -s charalarm -u root -p charalarm -o ./schemaspy/output -vizjs
```


## DB変更

テーブルにカラムを追加する場合などは `charalarm.ddl` に Alter文などを日付と共に追記する。

```
alter table user add token VARCHAR(255) NULL;
```


