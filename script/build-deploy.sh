#!/bin/bash -eu

cwd=`dirname $0`
cd $cwd/../application

################################################################
# Build
################################################################
# alarm
GOOS=linux GOARCH=amd64 go build -o build/alarm_add handler/alarm_add/alarm_add.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_edit handler/alarm_edit/alarm_edit.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_delete handler/alarm_delete/alarm_delete.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_list handler/alarm_list/alarm_list.go

# batch
GOOS=linux GOARCH=amd64 go build -o build/batch handler/batch/batch.go

# chara
GOOS=linux GOARCH=amd64 go build -o build/chara_list handler/chara_list/chara_list.go

# healthcheck
GOOS=linux GOARCH=amd64 go build -o build/healthcheck handler/healthcheck/healthcheck.go

# push-token
GOOS=linux GOARCH=amd64 go build -o build/push_token_ios_push_add handler/push_token_ios_push_add/push_token_ios_push_add.go
GOOS=linux GOARCH=amd64 go build -o build/push_token_ios_voip_push_add handler/push_token_ios_voip_push_add/push_token_ios_voip_push_add.go

# user
GOOS=linux GOARCH=amd64 go build -o build/user_info handler/user_info/user_info.go
GOOS=linux GOARCH=amd64 go build -o build/user_signup handler/user_signup/user_signup.go
GOOS=linux GOARCH=amd64 go build -o build/user_withdraw handler/user_withdraw/user_withdraw.go


################################################################
# Archive
################################################################
cd ./build

# alarm
zip alarm_add.zip alarm_add
zip alarm_delete.zip alarm_delete
zip alarm_edit.zip alarm_edit
zip alarm_list.zip alarm_list

# batch
zip batch.zip batch

# chara
zip chara_list.zip chara_list

# healthcheck
zip healthcheck.zip healthcheck

# push-notification
zip push_token_ios_push_add.zip push_token_ios_push_add
zip push_token_ios_voip_push_add.zip push_token_ios_voip_push_add

# user
zip user_signup.zip user_signup
zip user_info.zip user_info
zip user_withdraw.zip user_withdraw


################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r
cd $cwd



################################################################
# Upload
################################################################
aws s3 sync application/build s3://$S3_BUCKET_NAME/$APPLICATION_VERSION --exact-timestamps --delete



################################################################
# Deploy
################################################################

# alarm
aws lambda update-function-code --function-name alarm-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_add.zip
aws lambda update-function-code --function-name alarm-delete-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_delete.zip
aws lambda update-function-code --function-name alarm-edit-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_edit.zip
aws lambda update-function-code --function-name alarm-list-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_list.zip

# batch

# chara

# healthcheck
aws lambda update-function-code --function-name healthcheck-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/healthcheck.zip

# push-notification
aws lambda update-function-code --function-name push-token-ios-push-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/push_token_ios_push_add.zip
aws lambda update-function-code --function-name push-token-ios-voip-push-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/push_token_ios_voip_push_add.zip

# user
aws lambda update-function-code --function-name user-signup-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_signup.zip
aws lambda update-function-code --function-name user-withdraw-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_withdraw.zip
aws lambda update-function-code --function-name user-info-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_info.zip

