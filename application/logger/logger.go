package logger

import (
	"encoding/json"
	"fmt"
)

func Log(fileName string, funcName string, line int, err error) {
	message := LogMessage{
		Level:   "error",
		Message: err.Error(),
	}

	jsonBytes, _ := json.Marshal(message)
	fmt.Println(string(jsonBytes))
}
