package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trouble-ticket-ms/src/models"
	"trouble-ticket-ms/src/services"
	"trouble-ticket-ms/src/utils"
)

type AuthController interface {
	SignIn(context *gin.Context)
	SignUp(context *gin.Context)
}

type authController struct {
	authService services.AuthService
}

// SignIn User
// @Summary Sign In
// @Tags Auth
// @Param   request body     models.Auth  true  "Login credentials"
// @Success 200 {object} models.AuthJwtPayload
// @Failure 500 {object} error
// @Router /auth/signIn [post]
func (auth *authController) SignIn(context *gin.Context) {
	var authM models.Auth

	if !utils.BindJSON(context, &authM) {
		return
	}

	authJwtPayload, err := auth.authService.SignIn(authM)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, authJwtPayload)
}

// SignUp User
// @Summary Sign Up
// @Description Done by System Management Personnel's
// @Tags Auth
// @Param   request body     models.SignUpDTO  true  "Signup credentials"
// @Success 200 {object} any
// @Failure 500 {object} error
// @Router /auth/signUp [post]
// @Security Bearer
func (auth *authController) SignUp(context *gin.Context) {
	var signupDto models.SignUpDTO

	if !utils.BindJSON(context, &signupDto) {
		return
	}

	newUser, err := auth.authService.SignUp(signupDto)

	if err != nil {
		context.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, newUser)
}

func NewAuthController(ts services.AuthService) AuthController {
	return &authController{ts}
}
