package entity

type AnonymousAddPushTokenRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
	PushToken string `json:"pushToken"`
}

type AnonymousUserRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
}

type AnonymousDeleteAlarmRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
	AlarmID   string `json:"alarmID"`
}

type AnonymousAddAlarmRequest struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
	Alarm     Alarm  `json:"alarm"`
}
