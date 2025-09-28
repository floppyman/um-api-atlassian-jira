package api

import (
	"encoding/json"
	"fmt"
	"time"
	
	"github.com/floppyman/um-common/logging/logr"
)

// GetWorkLogIds calls the JIRA API (HTTP/HTTPS) with a unix timestamp.
// Returns a list up to 1000 entries or an error.
func GetWorkLogIds(start time.Time) (*WorklogUpdated, error) {
	var jsonObj *WorklogUpdated
	
	jsonArr, err := doGetRequest(fmt.Sprintf("worklog/updated?since=%d", start.UnixMilli()))
	if err != nil {
		return jsonObj, err
	}
	
	jsonErr := json.Unmarshal(jsonArr, &jsonObj)
	if jsonErr != nil {
		return jsonObj, jsonErr
	}
	
	return jsonObj, nil
}

// GetWorkLogItems calls the JIRA API (HTTP/HTTPS) with a list of up to 1000 Ids of Jira Worklogs.
// Return the list of worklog objects or an error.
func GetWorkLogItems(ids []int32) ([]WorklogItem, error) {
	var jsonObj []WorklogItem
	
	obj := WorklogForIds{
		Ids: ids,
	}
	
	j, err := json.Marshal(obj)
	if err != nil {
		return jsonObj, err
	}
	
	jsonArr, err := doPostRequest("worklog/list", j)
	if err != nil {
		return jsonObj, err
	}
	
	jsonErr := json.Unmarshal(jsonArr, &jsonObj)
	if jsonErr != nil {
		return jsonObj, jsonErr
	}
	
	return jsonObj, nil
}

func AddWorklog(issueIdOrKey string, startTime time.Time, timeSpent int32, comment string) (WorklogItem, error) {
	var jsonObj WorklogItem
	
	obj := WorklogAdd{
		Started:          startTime.Format("2006-01-02T15:04:05.000+0000"),
		TimeSpentSeconds: timeSpent,
		Comment: WorklogAddComment{
			Content: []WorklogAddCommentContent{
				{
					Content: []WorklogAddCommentContentContent{
						{
							Text: comment,
							Type: "text",
						},
					},
					Type: "paragraph",
				},
			},
			Type:    "doc",
			Version: 1,
		},
	}
	
	j, err := json.Marshal(obj)
	if err != nil {
		return jsonObj, err
	}
	
	jsonArr, err := doPostRequest(fmt.Sprintf("issue/%s/worklog", issueIdOrKey), j)
	if err != nil {
		return jsonObj, err
	}
	
	if doLogging {
		logr.Console.Trace().Msg(string(jsonArr))
	}
	
	jsonErr := json.Unmarshal(jsonArr, &jsonObj)
	if jsonErr != nil {
		return jsonObj, jsonErr
	}
	
	return jsonObj, nil
	
}
