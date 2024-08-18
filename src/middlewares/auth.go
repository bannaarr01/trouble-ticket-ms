package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

// AuthGuard middleware
func AuthGuard(deps services.AppDependencies) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get Token
		accessToken, err := utils.ExtractAuthTokenFromHeader(ctx)

		if err != nil {
			ctx.Abort()
			return
		}

		// invalid or expired token will fail and not be parsed
		_, parsedTokenClaim, err := deps.KeycloakClient.
			DecodeAccessToken(
				deps.Context,
				*accessToken,
				deps.KeycloakCfg.Realm,
			)

		if err != nil {
			log.Println("parse token failed:", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		userClaims, err := models.GetClaims(*parsedTokenClaim)

		if err != nil {
			log.Println("invalid user claims")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		// set auth user data to gin context
		ctx.Set("user", userClaims)
		ctx.Set("accessToken", *accessToken)

		// to the next request handler
		ctx.Next()
	}
}

// RoleGuard middleware
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, authUser := ctx.Get("user")

		if !authUser {
			log.Println("get user info failed:")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		userRoles := user.(*models.Claims).RealmAccess.Roles

		// Check if the user has any of the allowed roles
		hasAllowedRole := false
		for _, role := range userRoles {
			if utils.Contains(allowedRoles, role) {
				hasAllowedRole = true
				break
			}
		}

		if !hasAllowedRole {
			log.Println("You don't have the required roles to access resource")
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		// proceed
		ctx.Next()
	}
}
