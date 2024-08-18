package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJSON(context *gin.Context, target interface{}) bool {
	err := context.ShouldBindJSON(target)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		context.Abort()
		return false
	}

	return true
}
