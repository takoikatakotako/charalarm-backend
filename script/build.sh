#!/bin/bash -eu

cwd=`dirname $0`
cd $cwd/../application

################################################################
# Build
################################################################
# alarm
GOOS=linux GOARCH=amd64 go build -o build/alarm_add handler/alarm_add/alarm_add.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_delete handler/alarm_delete/alarm_delete.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_list handler/alarm_list/alarm_list.go

# batch
GOOS=linux GOARCH=amd64 go build -o build/batch handler/batch/batch.go

# chara
GOOS=linux GOARCH=amd64 go build -o build/chara_list handler/chara_list/chara_list.go

# healthcheck
GOOS=linux GOARCH=amd64 go build -o build/healthcheck handler/healthcheck/healthcheck.go

# user
GOOS=linux GOARCH=amd64 go build -o build/user_info_anonymous_user handler/user_info_anonymous_user/user_info_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/user_signup_anonymous_user handler/user_signup_anonymous_user/user_signup_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/user_withdraw_anonymous_user handler/user_withdraw_anonymous_user/user_withdraw_anonymous_user.go


################################################################
# Archive
################################################################
cd ./build

# alarm
zip alarm_add.zip alarm_add
zip alarm_delete.zip alarm_delete
zip alarm_list.zip alarm_list

# batch
zip batch.zip batch

# chara
zip chara_list.zip chara_list

# healthcheck
zip healthcheck.zip healthcheck

# user
zip user_signup_anonymous_user.zip user_signup_anonymous_user
zip user_info_anonymous_user.zip user_info_anonymous_user
zip user_withdraw_anonymous_user.zip user_withdraw_anonymous_user


################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r
cd $cwd
