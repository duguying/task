package utils

import (
	"encoding/json"
	// "fmt"
)

func JsonDecode(data string) interface{} {
	dataByte := []byte(data)
	var dat interface{}

	if err := json.Unmarshal(dataByte, &dat); err != nil {
		panic(err)
	}
	return dat
}

func JsonEncode(data interface{}) string {
	a, _ := json.Marshal(data)
	return string(a)
}
