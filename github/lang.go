package github

import (
	// "fmt"
	"task/utils"
)

var langDataTotal map[string]interface{}

func GetRepos() string {
	langDataTotal = make(map[string]interface{})

	jsonRepos, _ := utils.Get("https://api.github.com/users/duguying/repos")
	repos := utils.JsonDecode(jsonRepos)

	for _, v := range repos.([]interface{}) {
		reposName := v.(map[string]interface{})["name"].(string)
		getLangOfRepos(reposName)
	}

	ichartData := utils.Convert2Ichart(langDataTotal)
	return utils.JsonEncode(ichartData)
}

func getLangOfRepos(reposName string) {
	jsonLangData, _ := utils.Get("https://api.github.com/repos/duguying/" + reposName + "/languages")
	langData := utils.JsonDecode(jsonLangData)

	for k, v := range langData.(map[string]interface{}) {
		if _, ok := langDataTotal[k]; ok {
			langDataTotal[k] = langDataTotal[k].(float64) + v.(float64)
		} else {
			langDataTotal[k] = v.(float64)
		}
	}
}
