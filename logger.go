package logger

import (
	"fmt"
	"time"
)

type logger struct {
	Level     string                  `json:"level"`
	Timestamp string                  `json:"@timestamp"`
	Version   string                  `json:"@version"`
	App       string                  `json:"app"`
	Line      string                  `json:"line"`
	Method    string                  `json:"method"`
	Message   *map[string]interface{} `json:"message"`
}

var Log = &logger{
	Level: "",
	Timestamp: fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second()),
	Version: "1",
	App:     "app",
	Line:    "",
	Method:  "",
}
