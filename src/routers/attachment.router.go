package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type AttachmentRouter interface {
	SetAppRouting(*gin.Engine, services.AppDependencies)
}

type attachmentRouter struct {
	attachmentController controllers.AttachmentController
	deps                 services.AppDependencies
}

func NewAttachmentRouter(
	attachmentController controllers.AttachmentController,
	deps services.AppDependencies,
) AttachmentRouter {
	return &attachmentRouter{attachmentController, deps}
}

func (aRtr *attachmentRouter) SetAppRouting(server *gin.Engine, deps services.AppDependencies) {
	allowedRoles := []string{"super_admin", "admin", "assigner", "customer", "initiator"}

	v1 := server.Group("/api/v1")
	{
		attachments := v1.Group("/attachments").
			Use(middlewares.AuthGuard(deps), middlewares.RoleGuard(allowedRoles...))
		{
			attachments.POST(":ticketId", aRtr.attachmentController.Upload)
		}
	}

}
