package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
)

type AppRouter interface {
	SetAppRouting(server *gin.Engine)
}

type appRouter struct {
	appController controllers.AppController
}

func NewAppRouter(appController controllers.AppController) AppRouter {
	return &appRouter{appController}
}

func (appRtr *appRouter) SetAppRouting(server *gin.Engine) {
	server.GET("/api/v1", appRtr.appController.Index)
}
