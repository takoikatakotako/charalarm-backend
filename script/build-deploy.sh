#!/bin/bash

set -eu

cwd=`dirname $0`
cd $cwd/../application

pwd

################################################################
# Build
################################################################
# alarm
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/alarm_add handler/alarm_add/alarm_add.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/alarm_edit handler/alarm_edit/alarm_edit.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/alarm_delete handler/alarm_delete/alarm_delete.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/alarm_list handler/alarm_list/alarm_list.go

# call_batch
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/batch handler/call_batch/call_batch.go

# chara
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/chara_id handler/chara_id/chara_id.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/chara_list handler/chara_list/chara_list.go

# maintenance
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/maintenance handler/maintenance/maintenance.go

# healthcheck
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/healthcheck handler/healthcheck/healthcheck.go

# push-token
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/push_token_ios_push_add handler/push_token_ios_push_add/push_token_ios_push_add.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/push_token_ios_voip_push_add handler/push_token_ios_voip_push_add/push_token_ios_voip_push_add.go

# require
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/require handler/require/require.go

# user
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user_info handler/user_info/user_info.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user_signup handler/user_signup/user_signup.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user_withdraw handler/user_withdraw/user_withdraw.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/user_update_premium_plan handler/user_update_premium_plan/user_update_premium_plan.go

# call_worker
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/worker handler/call_worker/call_worker.go


################################################################
# Archive
################################################################
cd ./build

# alarm
zip alarm_add.zip alarm_add
zip alarm_delete.zip alarm_delete
zip alarm_edit.zip alarm_edit
zip alarm_list.zip alarm_list

# call_batch
zip batch.zip batch

# chara
zip chara_id.zip chara_id
zip chara_list.zip chara_list

# maintenance
zip maintenance.zip maintenance

# healthcheck
zip healthcheck.zip healthcheck

# push-notification
zip push_token_ios_push_add.zip push_token_ios_push_add
zip push_token_ios_voip_push_add.zip push_token_ios_voip_push_add

# require
zip require.zip require

# user
zip user_signup.zip user_signup
zip user_info.zip user_info
zip user_withdraw.zip user_withdraw

# call_worker
zip worker.zip worker

################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r



################################################################
# Upload
################################################################
aws s3 sync . s3://$S3_BUCKET_NAME/$APPLICATION_VERSION --exact-timestamps --delete



################################################################
# Deploy
################################################################

# alarm
aws lambda update-function-code --function-name alarm-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_add.zip
aws lambda update-function-code --function-name alarm-delete-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_delete.zip
aws lambda update-function-code --function-name alarm-edit-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_edit.zip
aws lambda update-function-code --function-name alarm-list-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/alarm_list.zip

# call_batch
aws lambda update-function-code --function-name batch-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/batch.zip

# chara
aws lambda update-function-code --function-name chara-id-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/chara_id.zip
aws lambda update-function-code --function-name chara-list-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/chara_list.zip

# maintenance
aws lambda update-function-code --function-name maintenance-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/maintenance.zip

# healthcheck
aws lambda update-function-code --function-name healthcheck-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/healthcheck.zip

# push-notification
aws lambda update-function-code --function-name push-token-ios-push-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/push_token_ios_push_add.zip
aws lambda update-function-code --function-name push-token-ios-voip-push-add-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/push_token_ios_voip_push_add.zip

# require
aws lambda update-function-code --function-name require-get-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/require.zip

# user
aws lambda update-function-code --function-name user-signup-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_signup.zip
aws lambda update-function-code --function-name user-withdraw-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_withdraw.zip
aws lambda update-function-code --function-name user-info-post-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/user_info.zip

# call_worker
aws lambda update-function-code --function-name worker-function --s3-bucket $S3_BUCKET_NAME --s3-key $APPLICATION_VERSION/worker.zip
