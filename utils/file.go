package utils

import (
	"io/ioutil"
)

func WriteString(path string, data string) error {
	var d1 = []byte(data)
	return ioutil.WriteFile(path, d1, 0666)
}
