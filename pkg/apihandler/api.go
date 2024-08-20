package apihandler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Run()  {
	router := gin.Default()
	
	router.GET("/", healthCheck)
	router.GET("/rc/validate/:manifest", rcValidate)
	router.POST("/jira/close", jiraClose)

	router.Run("0.0.0.0:8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "Server is healthy",
	})
}