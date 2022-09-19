#!/bin/bash -eu

################################################################
# Build
################################################################
GOOS=linux GOARCH=amd64 go build -o build/healthcheck healthcheck.go
GOOS=linux GOARCH=amd64 go build -o build/signup_anonymous_user signup_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/info_anonymous_user info_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/withdraw_anonymous_user withdraw_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_add alarm_add.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_delete alarm_delete.go
GOOS=linux GOARCH=amd64 go build -o build/alarm_list alarm_list.go

################################################################
# Archive
################################################################
cd ./build
zip healthcheck.zip healthcheck
zip signup_anonymous_user.zip signup_anonymous_user
zip info_anonymous_user.zip info_anonymous_user
zip withdraw_anonymous_user.zip withdraw_anonymous_user
zip alarm_add.zip alarm_add
zip alarm_delete.zip alarm_delete
zip alarm_list.zip alarm_list

################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r
cd ..
