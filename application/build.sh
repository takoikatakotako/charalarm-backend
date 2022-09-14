# healthcheck
GOOS=linux GOARCH=amd64 go build -o ./build/healthcheck healthcheck.go

# signup_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/signup_anonymous_user signup_anonymous_user.go

# info_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/info_anonymous_user info_anonymous_user.go

# withdraw_anonymous_user
GOOS=linux GOARCH=amd64 go build -o ./build/withdraw_anonymous_user withdraw_anonymous_user.go
