################################################################
# Build
################################################################
GOOS=linux GOARCH=amd64 go build -o build/healthcheck healthcheck.go
GOOS=linux GOARCH=amd64 go build -o build/signup_anonymous_user signup_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/info_anonymous_user info_anonymous_user.go
GOOS=linux GOARCH=amd64 go build -o build/withdraw_anonymous_user withdraw_anonymous_user.go

################################################################
# Archive
################################################################
cd ./build
zip healthcheck.zip healthcheck
zip signup_anonymous_user.zip signup_anonymous_user
zip info_anonymous_user.zip info_anonymous_user
zip withdraw_anonymous_user.zip withdraw_anonymous_user

################################################################
# Clear
################################################################
ls | grep -v -E '.zip$' | xargs rm -r
cd ..
