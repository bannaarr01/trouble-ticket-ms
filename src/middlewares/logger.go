package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Log middleware to format logs to file
func Log(appLog, errorLog *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		//2024/08/19 03:50:23 [2024/08/19 03:50:23] [GIN] ::1 | 200 | 53.8765ms | POST | /api/v1/auth/signIn
		appLog.Printf("[%s] [GIN] %s | %d | %v | %s | %s %s\n",
			end.Format("2006/01/02 15:04:05"),
			c.ClientIP(),
			c.Writer.Status(),
			latency,
			c.Request.Method,
			path,
			query,
		)

		if c.Writer.Status() >= 400 {
			errorLog.Printf("[%s] [GIN] %s | %d | %v | %s | %s %s - %s\n",
				end.Format("2006/01/02 15:04:05"),
				c.ClientIP(),
				c.Writer.Status(),
				latency,
				c.Request.Method,
				path,
				query,
				c.Errors.ByType(gin.ErrorTypePrivate).String(),
			)
		}
	}
}
