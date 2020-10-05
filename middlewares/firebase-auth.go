package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/marian0/expenses-go-api/common"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		common.RespondError(c, http.StatusBadRequest, "Empty token", "Authentication went wrong, please try again.")
		c.Abort()
		return
	}

	log.Printf(idToken)

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)

	if err != nil {
		common.RespondError(c, http.StatusBadRequest, err.Error(), "Authentication went wrong, please try again.")
		c.Abort()
		return
	}

	c.Set(common.CONST_UUID_IDENTIFIER, token.UID)
	c.Next()
}
