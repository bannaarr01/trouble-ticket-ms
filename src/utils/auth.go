package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ExtractAuthTokenFromHeader(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		// stop and no other code on the server runs
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no authorization header provided"})
	}

	extractedToken := strings.Split(token, "Bearer ")

	if extractedToken[1] == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization Bearer <token> is required"})
	}

	return extractedToken[1]
}
