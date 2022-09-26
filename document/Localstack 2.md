## LocalStack

### S3

```
aws s3 ls --endpoint-url=http://localhost:4572
```

### SQS

```
aws sns create-topic \
  --name my-topic \
  --endpoint-url http://localhost:4575
```

```
aws sns create-platform-application \
  --name ios-push-platform-application \
  --platform APNS \
  --attributes PlatformCredential=DAMMY \
  --endpoint-url http://localhost:4575
```

```
aws sns create-platform-endpoint \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-push-platform-application \
  --token MY_TOKEN \
  --endpoint-url http://localhost:4575
```

```
aws sns list-endpoints-by-platform-application \

```

```
aws sns   delete-platform-application \
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/my-topic3 \
  --endpoint-url http://localhost:4575
```

```
aws sns list-platform-applications \
  --endpoint-url http://localhost:4575
```

```
aws sns list-topics \
  --endpoint-url http://localhost:4575
```
  --platform-application-arn arn:aws:sns:ap-northeast-1:000000000000:app/APNS/ios-push-platform-application \
  --endpoint-url http://localhost:4575