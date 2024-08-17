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
	troubleTicketPrefix := server.Group("/api/v1/troubleTickets")
	//troubleTicketPrefix.POST("", tRtr.troubleTicketController.Create)
	troubleTicketPrefix.GET("", tRtr.troubleTicketController.FindAll)
	//troubleTicketPrefix.GET("/:id", tRtr.troubleTicketController.FindOne)
	//troubleTicketPrefix.PATCH("/:id", tRtr.troubleTicketController.Update)
	//troubleTicketPrefix.DELETE("/:id", tRtr.troubleTicketController.Remove)
}
