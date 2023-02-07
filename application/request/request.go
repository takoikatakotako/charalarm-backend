package request

type DeleteAlarmRequest struct {
	AlarmID   string `json:"alarmID"`
}

type AddAlarmRequest struct {
	Alarm     Alarm  `json:"alarm"`
}
