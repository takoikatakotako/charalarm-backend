package entity

type AnonymousDeleteAlarmRequest struct {
	UserID    string `json: "userID"`
	UserToken string `json: "userToken"`
	AlarmID   string `json: "alarmID"`
}
