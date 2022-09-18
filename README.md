# CharalarmBackend


## GET:  /healthcheck

ヘルスチェックに使用するエンドポイントです。

```
$ curl https://api.sandbox.swiswiswift.com/healthcheck
```

## POST: /user/signup/anonymous

ユーザーの新規登録を行うエンドポイントです

```
$ curl -X POST https://api.sandbox.swiswiswift.com/user/signup/anonymous \
    -H 'Content-Type: application/json' \
    -d '{"userID":"20f0c1cd-9c2a-411a-878c-9bd0bb15dc35","userToken":"038a5e28-15ce-46b4-8f46-4934202faa85"}'
```



## POST: /user/withdraw/anonymous


```
$ curl -X POST https://api.sandbox.swiswiswift.com/user/withdraw/anonymous \
    -H 'Content-Type: application/json' \
    -d '{"userID":"20f0c1cd-9c2a-411a-878c-9bd0bb15dc35","userToken":"038a5e28-15ce-46b4-8f46-4934202faa85"}'
```

## POST: /user/info/anonymous

ユーザーの情報を取得するエンドポイントです。

```
$ curl -X POST https://api.sandbox.swiswiswift.com/user/info/anonymous \
    -H 'Content-Type: application/json' \
    -d '{"userID":"20f0c1cd-9c2a-411a-878c-9bd0bb15dc35","userToken":"038a5e28-15ce-46b4-8f46-4934202faa85"}'
```


## GET:  /alarm/list


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

