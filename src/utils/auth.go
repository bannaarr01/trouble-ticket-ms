package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ExtractAuthTokenFromHeader from request authorization header
func ExtractAuthTokenFromHeader(ctx *gin.Context) (*string, error) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no authorization header provided"})
		return nil, errors.New("no authorization header provided")
	}

	// handle both lower n upperCase
	if strings.HasPrefix(token, "Bearer ") || strings.HasPrefix(token, "bearer ") {
		extractedToken := strings.TrimPrefix(token, "Bearer ")
		extractedToken = strings.TrimPrefix(extractedToken, "bearer ")
		return &extractedToken, nil
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization Bearer <token> is required"})
	return nil, errors.New("invalid authorization header format")
}

// Contains Helper function to check if a slice contains a certain element
func Contains(slice []string, elem string) bool {
	for _, x := range slice {
		if x == elem {
			return true
		}
	}
	return false
}
