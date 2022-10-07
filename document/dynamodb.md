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
  "alarmType":"{VOICE_CALL_ALARM or NEWS_CALL_ALARM or CALENDER_CALL_ALARM}",
  "alarmEnable":"{Bool}",
  "alarmName":"{String}",
  "alarmHour":"{Int}",
  "alarmMinute":"{Int}",
  "alarmTime":"{String}",
  "charaID":"{String}",
  "charaName":"{String}",
  "voiceFileURL":"{String}",  
  "sunday":"{Bool}",
  "monday":"{Bool}",
  "tuesday":"{Bool}",
  "wednesday":"{Bool}",
  "thursday":"{Bool}",
  "friday":"{Bool}",
  "saturday":"{Bool}"
}
```

### alarmType

アラームの種類です。

- VOICE_CALL_ALARM

キャラクターから電話がかかってきて、事前録音されたボイスを再生します。

- NEWS_CALL_ALARM

キャラクターから電話がかかってきて、音声合成されたニュースを再生します。

- CALENDER_CALL_ALARM

キャラクターから電話がかかってきて、音声合成された今日の予定を再生します。


### alarmTime

インデックス用のフィールドです。
`("%02d-%02d", alarmHour, alarmMinute)` が入ります。


### charaID

キャラクターのIDです。
ランダムの場合は `RANDOM` が入ります。


### charaName

キャラクター名です。
ランダムの場合は `RANDOM` が入ります。


### voiceFileURL

ボイス用のURLが入ります。
ランダムの場合は `RANDOM` が入ります。


## chara-table

```
{
  "charaID":"{UUID}",
  "charaEnable":"{Bool}",
  "charaName":"{String}",
  "charaDescription":"{String}",
  "charaProfiles":[
    {
      "title":{String},
      "name":{String},
      "url":{String}
    }
  ]

}
```
¥