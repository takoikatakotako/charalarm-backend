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
zip alarm_edit.zip alarm_edit
zip alarm_delete.zip alarm_delete
zip alarm_list.zip alarm_list

# batch
zip batch.zip batch

# chara
zip chara_list.zip chara_list

# healthcheck
zip healthcheck.zip healthcheck

# user
zip user_signup.zip user_signup
zip user_info.zip user_info
zip user_withdraw.zip user_withdraw


################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r
cd $cwd
