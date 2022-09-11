go get -u github.com/aws/aws-lambda-go/lambda
go env -w GO111MODULE=off
GOOS=linux GOARCH=amd64 go build -o hello
zip handler.zip ./hello

go get -u github.com/aws/aws-sdk-go
go get -u github.com/aws/aws-sdk-go-v2

go get -u github.com/aws/aws-sdk-go-v2/config
go get -u github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
go get -u github.com/aws/aws-sdk-go-v2/service/dynamodb
go get -u github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression

go: downloading github.com/aws/aws-sdk-go v1.44.95
go: downloading github.com/jmespath/go-jmespath v0.4.0
go get: added github.com/aws/aws-sdk-go v1.44.95
go get: added github.com/go-sql-driver/mysql v1.5.0
go get: added golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2



https://developer.so-tech.co.jp/entry/2022/03/08/120000

https://haithai91.medium.com/how-to-setup-aws-config-connect-to-localstack-with-golang-bd71ac2dd9d2





GOOS=linux GOARCH=amd64 go build -o signup_anonymous_user signup_anonymous_user.go
zip signup_anonymous_user.zip ./signup_anonymous_user

lambdaのハンドラをsignup_anonymous_userに設定する。

curl -X POST -H "Content-Type: application/json" -d '{"userId": "12345", "userToken":"テストユーザー"}' https://uj37p62g2liwpa334yjfh4e53i0jepqe.lambda-url.ap-northeast-1.on.aws/

