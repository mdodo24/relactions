package apihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"relactions/pkg/relhandler"
)

func jiraClose(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manifest := req.Manifest
	jira := req.Jira
	res := relhandler.ExecuteJiraClose(jira, manifest)
	if res.Error != "" {
		c.IndentedJSON(http.StatusBadRequest, res)
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}