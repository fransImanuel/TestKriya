package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("id")
		c.Set("id", id)

		//If Request Method is GET, Immidiately return
		if c.Request.Method == "GET" {
			c.Next()
			return
		}

		response := struct {
			Error_message string `json:"error_message"`
			Error_key     string `json:"error_key"`
		}{
			Error_message: "failed",
			Error_key:     "Anda Bukan Admin",
		}

		if !checkRoleByID(id) {
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		c.Next()
	}
}
