package apihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"relactions/pkg/relhandler"
)

func rcValidate(c *gin.Context) {
	manifest := c.Param("manifest")
	res := relhandler.ExecuteRCValidate(manifest)
	if res.Error != "" {
		c.IndentedJSON(http.StatusBadRequest, res)
		return
	}
	c.IndentedJSON(http.StatusOK, res)
	return
}