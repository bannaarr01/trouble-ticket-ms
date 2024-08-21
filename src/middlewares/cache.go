package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"time"
)

// Cache middleware for caching using Redis
func Cache(redisClient *redis.Client, defaultTTL time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Construct a more unique cache key
		cacheKey := fmt.Sprintf("%s-%s?%s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.URL.RawQuery)

		// Try to get data from Redis
		data, err := redisClient.Get(cacheKey).Result()
		if err == nil {
			// Data found in Redis, unmarshal it and return
			var cachedData interface{}
			err = json.Unmarshal([]byte(data), &cachedData)
			if err == nil {
				log.Printf("Cache hit: %s", cacheKey)
				ctx.JSON(http.StatusOK, cachedData)
				ctx.Abort()
				return
			}
			log.Printf("Failed to unmarshal cached data: %v", err)
		} else if err != redis.Nil {
			// An unexpected Redis error occurred
			log.Printf("Redis error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			ctx.Abort()
			return
		} // else // Cache miss

		// Proceed
		ctx.Next()

		// After handler execution, store response data in Redis right here
		value, exists := ctx.Get("data")
		if exists {
			resData, er := json.Marshal(value)
			if er != nil {
				log.Printf("Failed to marshal response data: %v", er)
				return
			}

			er = redisClient.Set(cacheKey, resData, defaultTTL).Err()
			if er != nil {
				log.Printf("Failed to set data in Redis: %v", er)
			}
			//else log.Printf("Cache set: %s", cacheKey)
		}
	}
}
