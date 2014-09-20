package utils

import (
// "fmt"
)

var config map[string]interface{}

func GetConfig(key string) interface{} {
	data := ReadFile("config.json")
	config = JsonDecode(data).(map[string]interface{})

	return config[key]
}
