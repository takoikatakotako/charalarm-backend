# DynamoDB

DynamoDBのテーブル構造についてです。


## user-table

ユーザー情報、認証情報、プッシュ通知用のトークンが格納されるテーブルです。

```
{
  "userID":"{UUID}",
  "userToken":"{UUID}",
  "iosVoIPPushTokens":{
    "token":"{iOSのVoIPプッシュ通知のトークン}",
    "snsEndpointArn":"{iOSのVoIPプッシュ通知のPlatformApplicationのエンドポイント}"
  },
  "iosPushTokens":{
    "token":"{iOSのVoIPプッシュ通知のトークン}",
    "snsEndpointArn":"{iOSのVoIPプッシュ通知のPlatformApplicationのエンドポイント}"
  }
}
```

## alarm-table

アラームの情報が入るテーブルです。

```
{
  "alarmID":"{UUID}",
  "userID":"{UUID}",
  "alarmType":"{XXXX or XXXXX or XXXXX or XXXX}",
  "alarmEnable":"{Bool}",
  "alarmName":"{String}",
  "alarmHour":"{Int}",
  "alarmMinute":"{Int}",
  "alarmTime":"{String}",
  "sunday":"{Bool}",
  "monday":"{Bool}",
  "tuesday":"{Bool}",
  "wednesday":"{Bool}",
  "thursday":"{Bool}",
  "friday":"{Bool}",
  "saturday":"{Bool}"
}
```
