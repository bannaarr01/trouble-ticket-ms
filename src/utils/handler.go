package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BindJSON attempts to bind the JSON body of a request to the provided target struct.
// It returns a boolean indicating whether the binding was successful.
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

// ParseID parses an ID from the Gin context and converts it to the specified integer type.
// It returns the parsed ID and an error. If parsing fails, it sets a 400 Bad Request response.
func ParseID[T ~int | ~int64 | ~uint | ~uint64](context *gin.Context, paramName string) (T, error) {
	idStr := context.Param(paramName)
	var id T

	// Parse the string to int64 first
	parsedID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return id, err
	}

	// Convert int64 to the target type
	id = T(parsedID)

	return id, nil
}

func ParseString(context *gin.Context, paramName string) (string, error) {
	paramValue := context.Param(paramName)
	if paramValue == "" {
		context.Error(errors.New("param not found"))
		context.JSON(http.StatusBadRequest, gin.H{"message": "param not found"})
		return "", errors.New("param not found")
	}
	return paramValue, nil
}
