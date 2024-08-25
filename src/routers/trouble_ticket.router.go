package routers

import (
	"github.com/gin-gonic/gin"
	"time"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type TroubleTicketRouter interface {
	SetAppRouting(*gin.Engine, services.AppDependencies)
}

type troubleTicketRouter struct {
	troubleTicketController controllers.TroubleTicketController
	deps                    services.AppDependencies
}

func NewTroubleTicketRouter(
	troubleTicketController controllers.TroubleTicketController,
	deps services.AppDependencies,
) TroubleTicketRouter {
	return &troubleTicketRouter{troubleTicketController, deps}
}

func (tRtr *troubleTicketRouter) SetAppRouting(server *gin.Engine, deps services.AppDependencies) {
	allowedRoles := []string{"super_admin", "admin", "assigner", "customer", "initiator"}

	v1 := server.Group("/api/v1")
	{ // method chaining/fluent interface to returns a new instance of gin.RouterGroup

		troubleTickets := v1.Group("/troubleTickets").
			Use(middlewares.AuthGuard(deps), middlewares.RoleGuard(allowedRoles...))
		{
			cached := troubleTickets.Use(middlewares.Cache(deps.RedisClient, 1*time.Hour))

			cached.GET("/filters", tRtr.troubleTicketController.FindAllFilter)
			troubleTickets.POST("", tRtr.troubleTicketController.Create)
			//show all if admin only, show only mine if not admin
			troubleTickets.GET("", tRtr.troubleTicketController.FindAll)
			troubleTickets.GET(":id", tRtr.troubleTicketController.FindOne)
			troubleTickets.PATCH(":id", tRtr.troubleTicketController.Update)
			troubleTickets.DELETE(":id", tRtr.troubleTicketController.Remove)
		}
	}

}
