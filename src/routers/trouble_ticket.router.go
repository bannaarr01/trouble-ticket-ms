package routers

import (
	"github.com/gin-gonic/gin"
	"trouble-ticket-ms/src/controllers"
)

type TroubleTicketRouter interface {
	SetAppRouting(context *gin.Engine)
}

type troubleTicketRouter struct {
	troubleTicketController controllers.TroubleTicketController
}

func NewTroubleTicketRouter(troubleTicketController controllers.TroubleTicketController) TroubleTicketRouter {
	return &troubleTicketRouter{troubleTicketController}
}

func (tRtr *troubleTicketRouter) SetAppRouting(server *gin.Engine) {
	v1 := server.Group("/api/v1")
	{ // method chaining/fluent interface to returns a new instance of gin.RouterGroup

		troubleTickets := v1.Group("/troubleTickets")
		{
			troubleTickets.GET("", tRtr.troubleTicketController.FindAll)
			//troubleTickets.POST("", tRtr.troubleTicketController.Create)
			//troubleTickets.GET("/:id", tRtr.troubleTicketController.FindOne)
			//troubleTickets.PATCH("/:id", tRtr.troubleTicketController.Update)
			//troubleTickets.DELETE("/:id", tRtr.troubleTicketController.Remove)
		}
	}

}
