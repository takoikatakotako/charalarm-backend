# DynamoDB

DynamoDBのテーブル構造についてです。


## user-table

ユーザー情報、認証情報、プッシュ通知用のトークンが格納されるテーブルです。

```
{
  "userID":"{UUID}",
  "authToken":"{UUID}",
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

キャラクターの情報が入るテーブルです。

```
{
  "charaID":"{String}",
  "enable":"{Bool}",
  "name":"{String}",
  "description":"{String}",
  "profiles":[
    {
      "title":{String},
      "name":{String},
      "url":{String}
    }
  ],
  "resources":{
    "images":[
      {String}
    ],
    "voices":[
      {String}
    ]
  },
  "expressions":{
    "{String}":{
      "images":[
        "{String}"
      ],
      "voices":[
        "{String}"
      ]
    }
  },
  "calls":{
    "voices":[
      "{String}"
    ]
  }
}
```

## サンプル

```
{
  "charaID":"com.charalarm.yui",
  "enable":true,
  "name":"井上結衣",
  "description":"井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！",
  "profiles":[
    {
      "title":"イラストレーター",
      "name":"さいもん",
      "url":"https://twitter.com/simon_ns"
    },
    {
      "title":"声優",
      "name":"Mai",
      "url":"https://twitter.com/mai_mizuiro"
    },
    {
      "title":"スクリプト",
      "name":"小旗ふたる！",
      "url":"https://twitter.com/Kass_kobataku"
    }
  ],
  "resources":{
    "images":[
      "thumbnail.png",
      "normal.png",
      "smile.png",
      "comfused.png"
    ],
    "voices":[
      "self-introduction.caf",
      "com-charalarm-yui-0.caf",
      "com-charalarm-yui-1.caf",
      "com-charalarm-yui-2.caf",
      "com-charalarm-yui-3.caf",
      "com-charalarm-yui-4.caf",
      "com-charalarm-yui-5.caf",
      "com-charalarm-yui-6.caf",
      "com-charalarm-yui-7.caf",
      "com-charalarm-yui-8.caf",
      "com-charalarm-yui-9.caf",
      "com-charalarm-yui-10.caf",
      "com-charalarm-yui-11.caf",
      "com-charalarm-yui-12.caf",
      "com-charalarm-yui-13.caf",
      "com-charalarm-yui-14.caf",
      "com-charalarm-yui-15.caf",
      "com-charalarm-yui-16.caf",
      "com-charalarm-yui-17.caf",
      "com-charalarm-yui-18.caf",
      "com-charalarm-yui-19.caf",
      "com-charalarm-yui-20.caf"
    ]
  },
  "expressions":{
    "normal":{
      "images":[
        "normal.png"
      ],
      "voices":[
        "com-charalarm-yui-1.caf",
        "com-charalarm-yui-4.caf",
        "com-charalarm-yui-5.caf"
      ]
    },
    "smile":{
      "images":[
        "smile.png"
      ],
      "voices":[
        "com-charalarm-yui-2.caf",
        "com-charalarm-yui-3.caf"
      ]
    },
    "comfused":{
      "images":[
        "comfused.png"
      ],
      "voices":[
        "com-charalarm-yui-5.caf",
        "com-charalarm-yui-12.caf",
        "com-charalarm-yui-13.caf",
        "com-charalarm-yui-14.caf"
      ]
    }
  },
  "calls":{
    "voices":[
      "com-charalarm-yui-15.caf",
      "com-charalarm-yui-16.caf",
      "com-charalarm-yui-17.caf",
      "com-charalarm-yui-18.caf",
      "com-charalarm-yui-19.caf",
      "com-charalarm-yui-20.caf"
    ]
  }
}
```

## CharaID

ドメインを逆にしたものが入ります。ex. com.charalarm.yui


```
aws dynamodb put-item \
--table-name chara-table \
--item '{"charaID":{"S":"com.charalarm.yui"},"charaEnable":{"BOOL":true},"charaName":{"S":"井上結衣"},"charaDescription":{"S":"井上結衣です。プログラマーとして働いていてこのアプリを作っています。このアプリをたくさん使ってくれると嬉しいです、よろしくね！"},"charaProfiles":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"charaResources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"},{"S":"smile.png"},{"S":"comfused.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"com-charalarm-yui-0.caf"},{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-6.caf"},{"S":"com-charalarm-yui-7.caf"},{"S":"com-charalarm-yui-8.caf"},{"S":"com-charalarm-yui-9.caf"},{"S":"com-charalarm-yui-10.caf"},{"S":"com-charalarm-yui-11.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"},{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"charaCall":{"M":{"voices":{"L":[{"S":"com-charalarm-yui-15.caf"},{"S":"com-charalarm-yui-16.caf"},{"S":"com-charalarm-yui-17.caf"},{"S":"com-charalarm-yui-18.caf"},{"S":"com-charalarm-yui-19.caf"},{"S":"com-charalarm-yui-20.caf"}]}}},"charaExpression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-1.caf"},{"S":"com-charalarm-yui-4.caf"},{"S":"com-charalarm-yui-5.caf"}]}}},"smile":{"M":{"images":{"L":[{"S":"smile.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-2.caf"},{"S":"com-charalarm-yui-3.caf"}]}}},"comfused":{"M":{"images":{"L":[{"S":"comfused.png"}]},"voices":{"L":[{"S":"com-charalarm-yui-5.caf"},{"S":"com-charalarm-yui-12.caf"},{"S":"com-charalarm-yui-13.caf"},{"S":"com-charalarm-yui-14.caf"}]}}}}}}' \
--profile sandbox
```


```
aws dynamodb put-item \
--table-name chara-table \
--item '{"charaID":{"S":"com.senpu-ki-soft.momiji"},"charaEnable":{"BOOL":true},"charaName":{"S":"紅葉"},"charaDescription":{"S":"金髪紅眼の美少女。疲れ気味のあなたを心配して様々な癒しを、と考えている。その正体は幾百年を生きる鬼の末裔。あるいはあなたに恋慕を抱く彼女。ちょっと素直になりきれないものの、なんやかんやいってそばにいてくれる面倒見のいい少女。日々あなたの生活を見届けている。「わっち？　名は紅葉でありんす。主様の支えになれるよう、掃除でもみみかきでもなんでも言っておくんなんし。か、かわいい？　い、いきなりそんなこと言わないでおくんなんし！」"},"charaProfiles":{"L":[{"M":{"title":{"S":"イラストレーター"},"name":{"S":"さいもん"},"url":{"S":"https://twitter.com/simon_ns"}}},{"M":{"title":{"S":"声優"},"name":{"S":"Mai"},"url":{"S":"https://twitter.com/mai_mizuiro"}}},{"M":{"title":{"S":"スクリプト"},"name":{"S":"小旗ふたる！"},"url":{"S":"https://twitter.com/Kass_kobataku"}}}]},"charaResources":{"M":{"images":{"L":[{"S":"thumbnail.png"},{"S":"normal.png"}]},"voices":{"L":[{"S":"self-introduction.caf"},{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"},{"S":"tap-head-1.caf"},{"S":"tap-head-2.caf"},{"S":"tap-head-3.caf"},{"S":"tap-lower-body-1.caf"},{"S":"tap-lower-body-2.caf"},{"S":"tap-lower-body-3.caf"},{"S":"tap-upper-body-1.caf"},{"S":"tap-upper-body-2.caf"},{"S":"tap-upper-body-3.caf"},{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"charaCall":{"M":{"voices":{"L":[{"S":"call-small-talk.caf"},{"S":"call-holiday-no-scheduled.caf"},{"S":"call-holiday-scheduled-alarm.caf"},{"S":"call-on-weekday-afternoon.caf"},{"S":"call-on-weekday-morning.caf"}]}}},"charaExpression":{"M":{"normal":{"M":{"images":{"L":[{"S":"normal.png"}]},"voices":{"L":[{"S":"tap-general-1.caf"},{"S":"tap-general-2.caf"},{"S":"tap-general-3.caf"},{"S":"tap-general-4.caf"},{"S":"tap-general-5.caf"}]}}}}}}' \
--profile sandbox
```
