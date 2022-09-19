package entity

type AnonymousAddAlarmRequest struct {
	UserID    string `json: "userID"`
	UserToken    string `json: "userToken"`
	Alarm Alarm `json: "alarm"`
}
