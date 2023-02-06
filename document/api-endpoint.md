
## GET: /healthcheck

ヘルスチェックに使用するエンドポイントです。

```
$ curl https://api.sandbox.swiswiswift.com/healthcheck | jq
```

```
{
  "message": "Healthy!"
}
```

## POST: /user/signup

ユーザーの新規登録を行うエンドポイントです。
`userID`, `userToken` はクライアント側で生成したUUIDを使用します。
生成したUUIDはクライアント側のKeyChainなどで保持します。

```
curl -X POST https://api.sandbox.swiswiswift.com/user/signup \
    -H 'Content-Type: application/json' \
    -d '{"userID":"20f0c1cd-9c2a-411a-878c-9bd0bb15dc35","userToken":"038a5e28-15ce-46b4-8f46-4934202faa85"}' | jq
```

成功時のレスポンスです。登録済みのユーザーでもこのレスポンスを返します。

```
{
  "message": "Sign Up Success!"
}
```

失敗時のレスポンスです。`userID`, `userToken` の形式がUUIDではない場合や予期せぬエラーが起きた場合はこのレスポンスを返します。

```
{
  "message": "Sign Up Failure..."
}
```


## POST: /user/withdraw

ユーザの退会時に使用するエンドポイントです。

```
BASIC_AUTH_HEADER=$(echo -n 20f0c1cd-9c2a-411a-878c-9bd0bb15dc35:038a5e28-15ce-46b4-8f46-4934202faa85 | base64)
curl -X POST https://api.sandbox.swiswiswift.com/user/withdraw \
    -H 'Content-Type: application/json' \
    -H "Authorization: Basic ${BASIC_AUTH_HEADER}"
```

```
{
  "message": "Withdraw Success!"
}
```

```
{
  "message": "Withdraw Failure..."
}
```


## POST: /user/info

ユーザーの情報を取得するエンドポイントです。

```
BASIC_AUTH_HEADER=$(echo -n 20f0c1cd-9c2a-411a-878c-9bd0bb15dc35:038a5e28-15ce-46b4-8f46-4934202faa85 | base64)
curl -X POST https://api.sandbox.swiswiswift.com/user/info \
    -H 'Content-Type: application/json' \
    -H "Authorization: Basic ${BASIC_AUTH_HEADER}" | jq
```

```
{
  "userID": "20f0c1cd-9c2a-411a-878c-9bd0bb15dc35",
  "userToken": "038a5e28-15ce-46b4-8f46-4934202faa85",
  "iosVoIPPushTokens": {
    "token": "",
    "snsEndpointArn": ""
  },
  "iosPushTokens": {
    "token": "",
    "snsEndpointArn": ""
  }
}
```


## POST:  /alarm/list

```
BASIC_AUTH_HEADER=$(echo -n 20f0c1cd-9c2a-411a-878c-9bd0bb15dc35:038a5e28-15ce-46b4-8f46-4934202faa85 | base64)
curl -X POST https://api.sandbox.swiswiswift.com/alarm/list \
    -H 'Content-Type: application/json' \
    -H "Authorization: Basic ${BASIC_AUTH_HEADER}" | jq
```

## POST: /alarm/add


## POST: /alarm/delete


## POST: /push-token/ios/push/add


## POST: /push-token/ios/voip-push/add


## GET: /news/list



## エラーメッセージ





curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"userId": "okki-", "userToken":"password"}' \
  https://api.sandbox.swiswiswift.com/user/signup/anonymous







https://99byleidca.execute-api.ap-northeast-1.amazonaws.com/production


https://api.sandbox.swiswiswift.com/user/signup/anonymous



aws lambda list-functions --profile sandbox | jq


aws lambda update-function-code --function-name healthcheck-get-function --s3-bucket application.charalarm.sandbox.swiswiswift.com --s3-key 0.0.1/healthcheck.zip --profile sandbox

