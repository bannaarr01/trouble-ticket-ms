package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type ExtIdentifierRouter interface {
	SetAppRouting(*gin.Engine, services.AppDependencies)
}

type extIdentifierRouter struct {
	extIdentifierController controllers.ExtIdentifierController
	deps                    services.AppDependencies
}

func NewExtIdentifierRouter(
	extIdentifierController controllers.ExtIdentifierController,
	deps services.AppDependencies,
) ExtIdentifierRouter {
	return &extIdentifierRouter{extIdentifierController, deps}
}

func (extRtr *extIdentifierRouter) SetAppRouting(server *gin.Engine, deps services.AppDependencies) {
	allowedRoles := []string{"super_admin", "admin", "assigner", "customer", "initiator"}

	v1 := server.Group("/api/v1")
	{
		attachments := v1.Group("/externalIdentifiers").
			Use(middlewares.AuthGuard(deps), middlewares.RoleGuard(allowedRoles...))
		{
			attachments.POST("/ticket/:id", extRtr.extIdentifierController.Create)
			attachments.GET("/ticket/:id", extRtr.extIdentifierController.FindByTicket)
			attachments.DELETE(":id", extRtr.extIdentifierController.Remove)
			attachments.GET(":id", extRtr.extIdentifierController.FindOne)
		}
	}

}
