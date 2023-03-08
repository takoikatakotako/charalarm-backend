package converter

import (
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/response"
)

func DatabaseUserToResponseUserInfo(user database.User) response.UserInfoResponse {
	return response.UserInfoResponse{
		UserID:           user.UserID,
		UserToken:        maskUserToken(user.AuthToken),
		IOSPushToken:     DatabasePushTokenToResponsePushToken(user.IOSVoIPPushToken),
		IOSVoIPPushToken: DatabasePushTokenToResponsePushToken(user.IOSVoIPPushToken),
	}
}

func DatabasePushTokenToResponsePushToken(pushToken database.PushToken) response.PushToken {
	return response.PushToken{
		Token:          pushToken.Token,
		SNSEndpointArn: pushToken.SNSEndpointArn,
	}
}

func EntityAlarmToDatabaseAlarm(alarm request.Alarm) database.Alarm {
	databaseAlarm := database.Alarm{
		ID:            alarm.AlarmID,
		UserID:        alarm.UserID,
		Type:          alarm.AlarmType,
		Enable:        alarm.AlarmEnable,
		Name:          alarm.AlarmName,
		Hour:          alarm.AlarmHour,
		Minute:        alarm.AlarmMinute,
		CharaID:       alarm.CharaID,
		CharaName:     alarm.CharaName,
		VoiceFilePath: alarm.VoiceFileURL,
		Sunday:        alarm.Sunday,
		Monday:        alarm.Monday,
		Tuesday:       alarm.Tuesday,
		Wednesday:     alarm.Wednesday,
		Thursday:      alarm.Thursday,
		Friday:        alarm.Friday,
		Saturday:      alarm.Saturday,
	}
	databaseAlarm.SetAlarmTime()
	return databaseAlarm
}

func DatabaseAlarmToEntityAlarm(alarm database.Alarm) request.Alarm {
	return request.Alarm{
		AlarmID:      alarm.ID,
		UserID:       alarm.UserID,
		AlarmType:    alarm.Type,
		AlarmEnable:  alarm.Enable,
		AlarmName:    alarm.Name,
		AlarmHour:    alarm.Hour,
		AlarmMinute:  alarm.Minute,
		CharaID:      alarm.CharaID,
		CharaName:    alarm.CharaName,
		VoiceFileURL: alarm.VoiceFilePath,
		Sunday:       alarm.Sunday,
		Monday:       alarm.Monday,
		Tuesday:      alarm.Tuesday,
		Wednesday:    alarm.Wednesday,
		Thursday:     alarm.Thursday,
		Friday:       alarm.Friday,
		Saturday:     alarm.Saturday,
	}
}

// 文字を*に変換
func maskUserToken(userToken string) string {
	length := len(userToken)
	var r string = ""
	for i := 0; i < length; i++ {
		if i == 0 {
			r += userToken[0:1]
		} else if i == 1 {
			r += userToken[1:2]
		} else {
			r += "*"
		}
	}
	return r
}
