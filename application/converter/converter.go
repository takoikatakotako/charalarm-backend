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
		AlarmID:       alarm.AlarmID,
		UserID:        alarm.UserID,
		Type:          alarm.Type,
		Enable:        alarm.Enable,
		Name:          alarm.Name,
		Hour:          alarm.Hour,
		Minute:        alarm.Minute,
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
		AlarmID:      alarm.AlarmID,
		UserID:       alarm.UserID,
		Type:         alarm.Type,
		Enable:       alarm.Enable,
		Name:         alarm.Name,
		Hour:         alarm.Hour,
		Minute:       alarm.Minute,
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
	return response.Alarm{
		AlarmID:      alarm.AlarmID,
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
