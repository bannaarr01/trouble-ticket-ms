package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"path/filepath"
	"trouble-ticket-ms/src/docs"
	"trouble-ticket-ms/src/logger"
	"trouble-ticket-ms/src/middlewares"
	"trouble-ticket-ms/src/services"
)

type MainRouter interface {
	setRouting(*gin.Engine, services.AppDependencies)
	StartServer(services.AppDependencies) error
}

type mainRouter struct {
	deps                services.AppDependencies
	appRouter           AppRouter
	authRouter          AuthRouter
	troubleTicketRouter TroubleTicketRouter
	attachmentRouter    AttachmentRouter
}

func NewMainRouter(
	deps services.AppDependencies,
	appRouter AppRouter,
	authRouter AuthRouter,
	troubleTicketRouter TroubleTicketRouter,
	attachmentRouter AttachmentRouter,
) MainRouter {
	return &mainRouter{
		deps,
		appRouter,
		authRouter,
		troubleTicketRouter,
		attachmentRouter,
	}
}

// setRouting sets up the routing for the application using the provided gin Engine.
// It delegates the routing setup to the router's instance.
func (mainRtr *mainRouter) setRouting(server *gin.Engine, deps services.AppDependencies) {
	mainRtr.appRouter.SetAppRouting(server)
	mainRtr.authRouter.SetAppRouting(server, deps)
	mainRtr.troubleTicketRouter.SetAppRouting(server, deps)
	mainRtr.attachmentRouter.SetAppRouting(server, deps)
}

// StartServer starts the server and returns an error if it fails.
// It sets up the routing using the setRouting method and then starts the server on the set port.
func (mainRtr *mainRouter) StartServer(deps services.AppDependencies) error {
	server := gin.Default()

	// cors
	server.Use(middlewares.CORS())

	// logger middleware
	appLog, errorLog := logger.NewLoggers()
	server.Use(middlewares.Log(appLog, errorLog))

	// Serve static files from project root dir "/data/*" under "/static/attachment/file/" path
	cwd, _ := os.Getwd()
	dirPath := filepath.Join(cwd, "data")
	server.Static("/static/attachment/file", dirPath)

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Trouble Ticket API"

	mainRtr.setRouting(server, deps)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return server.Run(":8080")
}
