package utils

import (
	// "fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}
