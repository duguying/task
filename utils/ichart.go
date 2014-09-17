package utils

import (
// "fmt"
// "math/rand"
)

func Convert2Ichart(data interface{}) interface{} {
	color := [15]string{"#FF8484", "#FFFF00", "#00FF00", "#00FFFF", "#0084FF", "#840000", "#8484C6", "#FF84FF", "#B44AB3", "#639B7F", "#C8B53B", "#BB5B47", "#C01CE1", "#A8F011", "#FFBA35"}
	length := len(data.(map[string]interface{}))
	objSlice := make([]interface{}, length)
	index := 0
	for k, v := range data.(map[string]interface{}) {
		objSlice[index] = interface{}(map[string]interface{}{"name": k, "value": v, "color": color[index%15]})
		index++
	}
	return objSlice
}
