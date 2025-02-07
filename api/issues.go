package api

import (
	"encoding/json"
)

// GetIssue calls the JIRA API (HTTP/HTTPS) to get the issue with the provided id.
// Returns the issue or an error
func GetIssue(issueId string) (IssueItem, error) {
	var jsonObj IssueItem

	// customfield_11000 - Activation (Yes, No)
	// customfield_11001 - Activation Comment
	// customfield_12212 - Activation Group

	jsonArr, err := doGetRequest("issue/" + issueId + "?fields=summary,key,customfield_10057,customfield_10058,customfield_10059")
	if err != nil {
		return jsonObj, err
	}

	jsonErr := json.Unmarshal(jsonArr, &jsonObj)
	if jsonErr != nil {
		return jsonObj, jsonErr
	}

	return jsonObj, nil
}
