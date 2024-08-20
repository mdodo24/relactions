package apihandler

type Request struct {
    Manifest string `json:"manifest"`
    Jira     string `json:"jira"`
}