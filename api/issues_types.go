package api

type IssueItem struct {
	Id     string          `json:"id"`
	Key    string          `json:"key"`
	Fields IssueItemFields `json:"fields"`
}

type IssueItemFields struct {
	Summary           string                         `json:"summary"`
	Activation        IssueItemFieldsActivation      `json:"customfield_10057"`
	ActivationComment string                         `json:"customfield_10058"`
	ActivationGroup   IssueItemFieldsActivationGroup `json:"customfield_10059"`
}

// customfield_10056 = Issue Payment
// customfield_10057 = Activation
// customfield_10058 = Activation Comment
// customfield_10059 = Activation Group

type IssueItemFieldsActivation struct {
	Value string `json:"value"`
}

type IssueItemFieldsActivationGroup struct {
	Value string `json:"value"`
}
