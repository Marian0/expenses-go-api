package common

import "github.com/gin-gonic/gin"

func respondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

//RespondError - Formats a json error message
func RespondError(c *gin.Context, status int, devError string, humanError string) {
	respondJSON(c, status, gin.H{"status": status, "devError": devError, "humanError": humanError})
}

//RespondSuccess - Formats a json success message
func RespondSuccess(c *gin.Context, status int, payload interface{}) {
	respondJSON(c, status, gin.H{"status": status, "data": payload})
}
