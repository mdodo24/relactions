package relhandler

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result   any `json:"result"`
	Error	string	`json:"error"`
}

type JiraCreateRespons struct {
	JiraID	string `json:"jiraID`
}