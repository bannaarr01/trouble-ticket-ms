package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
)

type AuthController interface {
	SignIn(context *gin.Context)
	SignOut(context *gin.Context)
}

type authController struct {
	authService services.AuthService
}

// SignIn User
// @Summary Sign In
// @Tags Auth
// @Param   request body     models.Auth  true  "Login credentials"
// @Success 200 {object} models.AuthTokenDTO
// @Failure 500 {object} error
// @Router /auth/signIn [post]
func (auth *authController) SignIn(context *gin.Context) {
	var authM models.Auth
	err := context.ShouldBindJSON(&authM)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	authLoginToken, err := auth.authService.SignIn(authM)

	if err != nil {
		log.Printf("error sign-in: %v", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, authLoginToken)
}

func (auth *authController) SignOut(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewAuthController(ts services.AuthService) AuthController {
	return &authController{ts}
}
