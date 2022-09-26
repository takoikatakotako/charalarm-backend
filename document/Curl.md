### 環境変数を設定

```
export ANONYMOUS_USER_NAME=9BF0D9AA-4C59-47D8-89E6-3BA6E66C4F7B
export ANONYMOUS_USER_PASSWORD=password
```

### ユーザー作成

```
curl "http://localhost:8080/api/anonymous/auth/signup" \
	-X POST \
	-H 'X-API-VERSION: 0' \
	-H 'Content-Type: application/json' \
	-d '{"anonymousUserName":"11893B3B-FB24-4B9C-AD7A-0034A707DD24","password":"9BF0D9AA-4C59-47D8-89E6-3BA6E66C4F7B"}'
```


### トークン登録

```
curl "http://localhost:8080/api/anonymous/token/ios/push/add" \
	-X POST \
	-H 'X-API-VERSION: 0' \
	-H 'Content-Type: application/json' \
	-d '{"anonymousUserName":"11893B3B-FB24-4B9C-AD7A-0034A707DD24","password":"9BF0D9AA-4C59-47D8-89E6-3BA6E66C4F7B","pushToken":"KABIKABI"}'
```

### キャラクター一覧表示

```
curl -H 'X-API-VERSION:0' http://localhost:8080/api/chara/list
```
