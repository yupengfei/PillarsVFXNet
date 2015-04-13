package utility

import (
	"encoding/json"
)

func ObjectToJsonString(object interface{}) *string {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	messageString := string(message)
	return &messageString
}

func ObjectToJsonByte(object interface{}) []byte {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}
