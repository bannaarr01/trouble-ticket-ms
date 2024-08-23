package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trouble-ticket-ms/src/utils"
)

// FileValidator Middleware validates file size and type
func FileValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, fileHeader, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if fileHeader.Size > 5<<20 { // 5MB
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "File size exceeds 5MB limit."})
			return
		}

		// Acceptable file types
		fileType := fileHeader.Header.Get("Content-Type")
		allowedTypes := []string{"image/png", "image/jpeg", "application/pdf", "application/msword"}
		if !utils.Contains(allowedTypes, fileType) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file type. Only png, jpg, pdf, and doc are allowed."})
			return
		}

		// update context
		ctx.Set("file", file)
		ctx.Set("fileHeader", fileHeader)

		ctx.Next()
	}
}
