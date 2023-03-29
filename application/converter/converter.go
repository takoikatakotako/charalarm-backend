package converter

import (
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/response"
)

func DatabaseUserToResponseUserInfo(user database.User) response.UserInfoResponse {
	return response.UserInfoResponse{
		UserID:           user.UserID,
		AuthToken:        maskAuthToken(user.AuthToken),
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

func RequestAlarmToDatabaseAlarm(alarm request.Alarm) database.Alarm {
	// request.Alarmは時差があるため、UTCのdatabase.Alarmに変換する
	var alarmHour int
	var alarmMinute int
	var alarmSunday bool
	var alarmMonday bool
	var alarmTuesday bool
	var alarmWednesday bool
	var alarmThursday bool
	var alarmFriday bool
	var alarmSaturday bool

	// 時差を計算
	diff := (float32(alarm.Hour) + float32(alarm.Minute)/60.0) - alarm.TimeDifference
	if diff > 24 {
		// tomorrow
		diff -= 24.0
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSunday = alarm.Monday
		alarmMonday = alarm.Tuesday
		alarmTuesday = alarm.Wednesday
		alarmWednesday = alarm.Thursday
		alarmThursday = alarm.Friday
		alarmFriday = alarm.Saturday
		alarmSaturday = alarm.Sunday
	} else if diff >= 0 {
		// today
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSunday = alarm.Sunday
		alarmMonday = alarm.Monday
		alarmTuesday = alarm.Tuesday
		alarmWednesday = alarm.Wednesday
		alarmThursday = alarm.Thursday
		alarmFriday = alarm.Friday
		alarmSaturday = alarm.Saturday
	} else {
		// yesterday
		diff += 24.0
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSunday = alarm.Saturday
		alarmMonday = alarm.Sunday
		alarmTuesday = alarm.Monday
		alarmWednesday = alarm.Tuesday
		alarmThursday = alarm.Wednesday
		alarmFriday = alarm.Thursday
		alarmSaturday = alarm.Friday
	}

	databaseAlarm := database.Alarm{
		AlarmID:        alarm.AlarmID,
		UserID:         alarm.UserID,
		Type:           alarm.Type,
		Enable:         alarm.Enable,
		Name:           alarm.Name,
		Hour:           alarmHour,
		Minute:         alarmMinute,
		TimeDifference: alarm.TimeDifference,
		CharaID:        alarm.CharaID,
		CharaName:      alarm.CharaName,
		VoiceFileName:  alarm.VoiceFileName,
		Sunday:         alarmSunday,
		Monday:         alarmMonday,
		Tuesday:        alarmTuesday,
		Wednesday:      alarmWednesday,
		Thursday:       alarmThursday,
		Friday:         alarmFriday,
		Saturday:       alarmSaturday,
	}
	databaseAlarm.SetAlarmTime()
	return databaseAlarm
}

func DatabaseCharaListToResponseCharaList(charaList []database.Chara) []response.Chara {
	responseCharaList := make([]response.Chara, 0)
	for i := 0; i < len(charaList); i++ {
		responseChara := DatabaseCharaToResponseChara(charaList[i])
		responseCharaList = append(responseCharaList, responseChara)
	}
	return responseCharaList
}

func DatabaseCharaToResponseChara(databaseChara database.Chara) response.Chara {
	return response.Chara{
		CharaID:     databaseChara.CharaID,
		Enable:      databaseChara.Enable,
		Name:        databaseChara.Name,
		Description: databaseChara.Description,
		Profiles:    databaseCharaProfileListToResponseCharaProfileList(databaseChara.CharaProfiles),
		Resources:   databaseCharaResourceListToResponseCharaResourceList(databaseChara.CharaResources),
		Expression:  databaseCharaExpressionMapToResponseCharaExpressionMap(databaseChara.CharaExpressions),
		Calls:       databaseCharaCallListToResponseCharaCallList(databaseChara.CharaCalls),
	}
}

func databaseCharaProfileListToResponseCharaProfileList(databaseCharaProfileList []database.CharaProfile) []response.CharaProfile {
	responseCharaProfileList := make([]response.CharaProfile, 0)
	for i := 0; i < len(databaseCharaProfileList); i++ {
		responseCharaProfile := databaseCharaProfileToResponseCharaProfile(databaseCharaProfileList[i])
		responseCharaProfileList = append(responseCharaProfileList, responseCharaProfile)
	}
	return responseCharaProfileList
}

func databaseCharaProfileToResponseCharaProfile(databaseCharaProfile database.CharaProfile) response.CharaProfile {
	return response.CharaProfile{
		Title: databaseCharaProfile.Title,
		Name:  databaseCharaProfile.Name,
		URL:   databaseCharaProfile.URL,
	}
}

func databaseCharaResourceListToResponseCharaResourceList(databaseCharaResourceList []database.CharaResource) []response.CharaResource {
	responseCharaResourceList := make([]response.CharaResource, 0)
	for i := 0; i < len(databaseCharaResourceList); i++ {
		responseCharaResource := databaseCharaResourceToResponseCharaResource(databaseCharaResourceList[i])
		responseCharaResourceList = append(responseCharaResourceList, responseCharaResource)
	}
	return responseCharaResourceList
}

func databaseCharaResourceToResponseCharaResource(databaseCharaResource database.CharaResource) response.CharaResource {
	return response.CharaResource{
		DirectoryName: databaseCharaResource.DirectoryName,
		FileName:      databaseCharaResource.FileName,
	}
}

func databaseCharaExpressionMapToResponseCharaExpressionMap(databaseCharaExpressionMap map[string]database.CharaExpression) map[string]response.CharaExpression {
	responseCharaExpressionMap := map[string]response.CharaExpression{}
	for k, v := range databaseCharaExpressionMap {
		responseCharaExpression := response.CharaExpression{
			Images: v.Images,
			Voices: v.Voices,
		}
		responseCharaExpressionMap[k] = responseCharaExpression
	}
	return responseCharaExpressionMap
}

func databaseCharaCallListToResponseCharaCallList(databaseCharaCallList []database.CharaCall) []response.CharaCall {
	responseCharaCallList := make([]response.CharaCall, 0)
	for i := 0; i < len(databaseCharaCallList); i++ {
		responseCharaCall := databaseCharaCallToResponseCharaCall(databaseCharaCallList[i])
		responseCharaCallList = append(responseCharaCallList, responseCharaCall)
	}
	return responseCharaCallList
}

func databaseCharaCallToResponseCharaCall(databaseCharaCall database.CharaCall) response.CharaCall {
	return response.CharaCall{
		Message: databaseCharaCall.Message,
		Voice:   databaseCharaCall.Voice,
	}
}

func DatabaseAlarmToResponseAlarm(alarm database.Alarm) response.Alarm {
	// UTCのdatabase.Alarmを時差のあるresponse.Alarmに変換する
	var alarmHour int
	var alarmMinute int
	var alarmSunday bool
	var alarmMonday bool
	var alarmTuesday bool
	var alarmWednesday bool
	var alarmThursday bool
	var alarmFriday bool
	var alarmSaturday bool

	// 時差を計算
	diff := (float32(alarm.Hour) + float32(alarm.Minute)/60.0) + alarm.TimeDifference
	if diff > 24 {
		// tomorrow
		diff -= 24.0
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSunday = alarm.Monday
		alarmMonday = alarm.Tuesday
		alarmTuesday = alarm.Wednesday
		alarmWednesday = alarm.Thursday
		alarmThursday = alarm.Friday
		alarmFriday = alarm.Saturday
		alarmSaturday = alarm.Sunday
	} else if diff >= 0 {
		// today
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSunday = alarm.Sunday
		alarmMonday = alarm.Monday
		alarmTuesday = alarm.Tuesday
		alarmWednesday = alarm.Wednesday
		alarmThursday = alarm.Thursday
		alarmFriday = alarm.Friday
		alarmSaturday = alarm.Saturday
	} else {
		// yesterday
		diff += 24.0
		alarmHour = int(diff)
		alarmMinute = int((diff - float32(alarmHour)) * 60)
		alarmSaturday = alarm.Friday
		alarmFriday = alarm.Thursday
		alarmThursday = alarm.Wednesday
		alarmWednesday = alarm.Tuesday
		alarmTuesday = alarm.Monday
		alarmMonday = alarm.Sunday
		alarmSunday = alarm.Saturday
	}

	return response.Alarm{
		AlarmID:        alarm.AlarmID,
		UserID:         alarm.UserID,
		Type:           alarm.Type,
		Enable:         alarm.Enable,
		Name:           alarm.Name,
		Hour:           alarmHour,
		Minute:         alarmMinute,
		TimeDifference: alarm.TimeDifference,
		CharaID:        alarm.CharaID,
		CharaName:      alarm.CharaName,
		VoiceFileName:  alarm.VoiceFileName,
		Sunday:         alarmSunday,
		Monday:         alarmMonday,
		Tuesday:        alarmTuesday,
		Wednesday:      alarmWednesday,
		Thursday:       alarmThursday,
		Friday:         alarmFriday,
		Saturday:       alarmSaturday,
	}
}

// 文字を*に変換
func maskAuthToken(authToken string) string {
	length := len(authToken)
	var r = ""
	for i := 0; i < length; i++ {
		if i == 0 {
			r += authToken[0:1]
		} else if i == 1 {
			r += authToken[1:2]
		} else {
			r += "*"
		}
	}
	return r
}
