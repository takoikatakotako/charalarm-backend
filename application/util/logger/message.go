package logger

type LogMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
