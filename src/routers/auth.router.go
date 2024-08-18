package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type AuthRouter interface {
	SetAppRouting(server *gin.Engine, deps services.AppDependencies)
}

type authRouter struct {
	authController controllers.AuthController
	deps           services.AppDependencies
}

func NewAuthRouter(authController controllers.AuthController, deps services.AppDependencies) AuthRouter {
	return &authRouter{authController, deps}
}

func (authRtr *authRouter) SetAppRouting(server *gin.Engine, deps services.AppDependencies) {
	// allowed roles to manage user
	allowedRoles := []string{"super_admin", "admin"}

	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			// No Authentication needed
			auth.POST("/signIn", authRtr.authController.SignIn)

			// Grouped system management access routes: must be Authenticated and authorized to access
			sysMgtAccessGroup := auth.Use(middlewares.AuthGuard(deps), middlewares.RoleGuard(allowedRoles...))
			sysMgtAccessGroup.POST("/signUp", authRtr.authController.SignUp)

		}
	}

}
