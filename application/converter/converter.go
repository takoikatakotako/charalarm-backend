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

func DatabaseCharaListToResponseCharaList(charaList []database.Chara) []response.Chara {
	var responseCharaList []response.Chara
	for i := 0; i < len(charaList); i++ {
		responseChara := databaseCharaToResponseChara(charaList[i])
		responseCharaList = append(responseCharaList, responseChara)
	}
	return responseCharaList
}

func databaseCharaToResponseChara(databaseChara database.Chara) response.Chara {
	return response.Chara{
		CharaID:          databaseChara.CharaID,
		CharaEnable:      databaseChara.CharaEnable,
		CharaName:        databaseChara.CharaName,
		CharaDescription: databaseChara.CharaDescription,
		CharaProfiles:    databaseCharaProfileListToResponseCharaProfileList(databaseChara.CharaProfiles),
		CharaResource:    databaseCharaResourceToResponseCharaResource(databaseChara.CharaResource),
		CharaExpression:  databaseCharaExpressionMapToResponseCharaExpressionMap(databaseChara.CharaExpression),
		CharaCall:        databaseCharaCallToResponseCharaCall(databaseChara.CharaCall),
	}
}

func databaseCharaProfileListToResponseCharaProfileList(databaseCharaProfileList []database.CharaProfile) []response.CharaProfile {
	var responseCharaList []response.CharaProfile
	for i := 0; i < len(databaseCharaProfileList); i++ {
		responseCharaProfile := databaseCharaProfileToResponseCharaProfile(databaseCharaProfileList[i])
		responseCharaList = append(responseCharaList, responseCharaProfile)
	}
	return responseCharaList
}

func databaseCharaProfileToResponseCharaProfile(databaseCharaProfile database.CharaProfile) response.CharaProfile {
	return response.CharaProfile{
		Title: databaseCharaProfile.Title,
		Name:  databaseCharaProfile.Name,
		URL:   databaseCharaProfile.URL,
	}
}

func databaseCharaResourceToResponseCharaResource(databaseCharaResource database.CharaResource) response.CharaResource {
	return response.CharaResource{
		Images: databaseCharaResource.Images,
		Voices: databaseCharaResource.Voices,
	}
}

func databaseCharaExpressionMapToResponseCharaExpressionMap(databaseCharaExpressionMap map[string]database.CharaExpression) map[string]response.CharaExpression {
	var responseCharaExpressionMap map[string]response.CharaExpression
	for k, v := range databaseCharaExpressionMap {
		responseCharaExpression := response.CharaExpression{
			Images: v.Images,
			Voices: v.Voices,
		}
		responseCharaExpressionMap[k] = responseCharaExpression
	}
	return responseCharaExpressionMap
}

func databaseCharaCallToResponseCharaCall(databaseCharaCall database.CharaCall) response.CharaCall {
	return response.CharaCall{
		Voices: databaseCharaCall.Voices,
	}
}

func DatabaseAlarmToResponseAlarm(alarm database.Alarm) response.Alarm {
	return response.Alarm{
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
func maskAuthToken(authToken string) string {
	length := len(authToken)
	var r string = ""
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
