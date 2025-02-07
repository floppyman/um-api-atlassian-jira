package api

type WorklogUpdated struct {
	Values   []WorklogUpdatedValue `json:"values"`
	Since    int64                 `json:"since"`
	Until    int64                 `json:"until"`
	Self     string                `json:"self"`
	NextPage string                `json:"nextPage"`
	LastPage bool                  `json:"lastPage"`
}

type WorklogUpdatedValue struct {
	WorklogId   int32 `json:"worklogId"`
	UpdatedTime int64 `json:"updatedTime"`
}

type WorklogForIds struct {
	Ids []int32 `json:"ids"`
}

type WorklogItem struct {
	Self             string            `json:"self"`
	Author           WorklogItemAuthor `json:"author"`
	IssueId          string            `json:"issueId"`
	TimeSpentSeconds int32             `json:"timeSpentSeconds"`
	Started          string            `json:"started"`
	Updated          string            `json:"updated"`
	Created          string            `json:"created"`
}

type WorklogItemAuthor struct {
	DisplayName  string `json:"displayName"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	AccountId    string `json:"accountId"`
}

type WorklogAdd struct {
	Comment          WorklogAddComment `json:"comment"`
	Started          string            `json:"started"`
	TimeSpentSeconds int32             `json:"timeSpentSeconds"`
}
type WorklogAddComment struct {
	Content []WorklogAddCommentContent `json:"content"`
	Type    string                     `json:"type"`
	Version int32                      `json:"version"`
}
type WorklogAddCommentContent struct {
	Content []WorklogAddCommentContentContent `json:"content"`
	Type    string                            `json:"type"`
}
type WorklogAddCommentContentContent struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
