package relhandler

import "fmt"

func ExecuteJiraCreate(manifest string) Response {
	if manifest == "" {
		return Response{
			Status: "not ok",
			Error: "empty manifest or jira",
		}
	}

	JiraCreateResponse := JiraCreateRespons{
		JiraID: "MA-1234",
	}

	return Response{
		Status: "ok",
		Result: JiraCreateResponse,
		Message: "jira create successfully",
	}
}

func ExecuteJiraClose(manifest, jira string) Response {
	if manifest == "" || jira == "" {
		return Response{
			Status: "not ok",
			Error: "empty manifest or jira",
		}
	}

	return Response{
		Status: "ok",
		Message: fmt.Sprintf("jira %v closed successfully", jira),
	}
}