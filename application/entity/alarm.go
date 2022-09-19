package entity

type Alarm struct {
	AlarmID   string `json: "alarmID"`
	AlarmType string `json: "alarmID"`
	Enable    bool   `json: "enable"`
	Name      bool   `json: "name"`
	Hour      int    `json: "hour"`
	Minute    int    `json: "minute"`
	Sunday    bool   `json: "sunday"`
	Monday    bool   `json: "monday"`
	Tuesday   bool   `json: "tuesday"`
	Wednesday bool   `json: "wednesday"`
	Thursday  bool   `json: "thursday"`
	Friday    bool   `json: "friday"`
	Saturday  bool   `json: "saturday"`
}
