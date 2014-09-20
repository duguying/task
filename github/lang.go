package github

import (
	"task/utils"
)

var langDataTotal map[string]interface{}

func GetRepos() string {
	langDataTotal = make(map[string]interface{})
	jsonRepos, _ := utils.Get("https://api.github.com/users/" + utils.GetConfig("user").(string) + "/repos?token=" + utils.GetConfig("token").(string))
	repos := utils.JsonDecode(jsonRepos)

	for _, v := range repos.([]interface{}) {
		reposName := v.(map[string]interface{})["name"].(string)
		getLangOfRepos(reposName)
	}

	ichartData := utils.Convert2Ichart(langDataTotal)
	return utils.JsonEncode(ichartData)
}

func getLangOfRepos(reposName string) {
	jsonLangData, _ := utils.Get("https://api.github.com/repos/" + utils.GetConfig("user").(string) + "/" + reposName + "/languages?token=" + utils.GetConfig("token").(string))
	langData := utils.JsonDecode(jsonLangData)

	for k, v := range langData.(map[string]interface{}) {
		if _, ok := langDataTotal[k]; ok {
			langDataTotal[k] = langDataTotal[k].(float64) + v.(float64)
		} else {
			langDataTotal[k] = v.(float64)
		}
	}
}
