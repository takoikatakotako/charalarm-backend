# healthcheck
GOOS=linux GOARCH=amd64 go build -o ./build/healthcheck healthcheck.go
zip ./build/healthcheck.zip ./build/healthcheck

# signup_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/signup_anonymous_user signup_anonymous_user.go
zip ./build/signup_anonymous_user.zip ./build/signup_anonymous_user

# info_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/info_anonymous_user info_anonymous_user.go
zip ./build/info_anonymous_user.zip ./build/info_anonymous_user

# withdraw_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/withdraw_anonymous_user withdraw_anonymous_user.go
zip ./build/withdraw_anonymous_user.zip ./build/withdraw_anonymous_user


# remove not zip file
cd build && ls | grep -v -E '.zip$' | xargs rm -r
