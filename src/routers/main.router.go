package routers

import (
	"github.com/gin-gonic/gin"
)

type MainRouter interface {
	setRouting(*gin.Engine)
	StartServer() error
}

type mainRouter struct {
	appRouter           AppRouter
	troubleTicketRouter TroubleTicketRouter
}

func NewMainRouter(appRouter AppRouter, troubleTicketRouter TroubleTicketRouter) MainRouter {
	return &mainRouter{
		appRouter,
		troubleTicketRouter,
	}
}

// setRouting sets up the routing for the application using the provided gin Engine.
// It delegates the routing setup to the router's instance.
func (mainRtr *mainRouter) setRouting(server *gin.Engine) {
	mainRtr.appRouter.SetAppRouting(server)
	mainRtr.troubleTicketRouter.SetAppRouting(server)
}

// StartServer starts the server and returns an error if it fails.
// It sets up the routing using the setRouting method and then starts the server on the set port.
func (mainRtr *mainRouter) StartServer() error {
	server := gin.Default()
	mainRtr.setRouting(server)

	return server.Run(":8080")
}
