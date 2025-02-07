package api

import "encoding/json"

// GetProjects calls the JIRA API (HTTP/HTTPS) to get a list of all projects.
// Return a list of all projects or an error.
func GetProjects() ([]ProjectItem, error) {
	var jsonObj []ProjectItem

	jsonArr, err := doGetRequest("project")
	if err != nil {
		return jsonObj, err
	}

	jsonErr := json.Unmarshal(jsonArr, &jsonObj)
	if jsonErr != nil {
		return jsonObj, jsonErr
	}

	return jsonObj, nil
}

// // PROJECT ITEM EXAMPLE
// --------------------------------------------------------
// {
// 	"expand": "description,lead,url,projectKeys",
// 	"self": "https://jira.dev.r2p.com/rest/api/2/project/11300",
// 	"id": "11300",
// 	"key": "FSUP",
// 	"name": "(INACTIVE)Flensburg - Support",
// 	"avatarUrls": {
// 		"48x48": "https://jira.dev.r2p.com/secure/projectavatar?avatarId=10324",
// 		"24x24": "https://jira.dev.r2p.com/secure/projectavatar?size=small&avatarId=10324",
// 		"16x16": "https://jira.dev.r2p.com/secure/projectavatar?size=xsmall&avatarId=10324",
// 		"32x32": "https://jira.dev.r2p.com/secure/projectavatar?size=medium&avatarId=10324"
// 	},
// 	"projectCategory": {
// 		"self": "https://jira.dev.r2p.com/rest/api/latest/projectCategory/10001",
// 		"id": "10001",
// 		"name": "Flensburg",
// 		"description": ""
// 	},
// 	"projectTypeKey": "business",
// 	"archived": false
// }
// --------------------------------------------------------
