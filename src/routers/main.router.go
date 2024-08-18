package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"trouble-ticket-ms/src/docs"
	"trouble-ticket-ms/src/logger"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type MainRouter interface {
	setRouting(*gin.Engine, services.AppDependencies)
	StartServer(deps services.AppDependencies) error
}

type mainRouter struct {
	deps                services.AppDependencies
	appRouter           AppRouter
	authRouter          AuthRouter
	troubleTicketRouter TroubleTicketRouter
}

func NewMainRouter(
	deps services.AppDependencies,
	appRouter AppRouter,
	authRouter AuthRouter,
	troubleTicketRouter TroubleTicketRouter,
) MainRouter {
	return &mainRouter{
		deps,
		appRouter,
		authRouter,
		troubleTicketRouter,
	}
}

// setRouting sets up the routing for the application using the provided gin Engine.
// It delegates the routing setup to the router's instance.
func (mainRtr *mainRouter) setRouting(server *gin.Engine, deps services.AppDependencies) {
	mainRtr.appRouter.SetAppRouting(server)
	mainRtr.authRouter.SetAppRouting(server, deps)
	mainRtr.troubleTicketRouter.SetAppRouting(server)
}

// StartServer starts the server and returns an error if it fails.
// It sets up the routing using the setRouting method and then starts the server on the set port.
func (mainRtr *mainRouter) StartServer(deps services.AppDependencies) error {
	server := gin.Default()

	// logger middleware
	appLog, errorLog := logger.NewLoggers()
	server.Use(middlewares.Log(appLog, errorLog))

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Trouble Ticket API"

	mainRtr.setRouting(server, deps)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server.Run(":8080")
}
