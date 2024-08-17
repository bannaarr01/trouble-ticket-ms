package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"trouble-ticket-ms/src/docs"
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

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Trouble Ticket API"

	mainRtr.setRouting(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server.Run(":8080")
}
