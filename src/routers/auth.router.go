package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
)

type AuthRouter interface {
	SetAppRouting(context *gin.Engine)
}

type authRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) AuthRouter {
	return &authRouter{authController}
}

func (authRtr *authRouter) SetAppRouting(server *gin.Engine) {
	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signIn", authRtr.authController.SignIn)
		}
	}

}
