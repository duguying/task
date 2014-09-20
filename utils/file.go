package utils

import (
	"io/ioutil"
	"os"
)

func WriteFile(path string, data string) error {
	var d1 = []byte(data)
	return ioutil.WriteFile(path, d1, 0666)
}

// 读取文本文件
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return string(fd)
}
