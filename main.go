package main

import (
	"fmt"
	"task/github"
	"task/utils"
)

func main() {

	json := github.GetRepos()
	fmt.Println(json)
	utils.WriteString("./data.json", json)

}
