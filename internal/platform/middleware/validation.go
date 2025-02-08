package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateJSONMiddleware(requiredFields ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request map[string]interface{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}

		for _, field := range requiredFields {
			if _, exists := request[field]; !exists {
				c.JSON(http.StatusBadRequest, gin.H{"error": field + " is required"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
