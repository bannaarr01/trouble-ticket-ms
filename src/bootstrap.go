package main

import (
	"log"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/routers"
	"trouble-ticket-ms/src/services"
)

func bootstrap(dbConn *db.DB) {
	// dep
	dependencies := services.InitAppDependencies()

	// repositories
	attachmentRepo := repositories.NewAttachmentRepository(dbConn)
	troubleTicketRepo := repositories.NewTroubleTicketRepository(dbConn)

	// services
	authService := services.NewAuthService(*dependencies)
	attachmentService := services.NewAttachmentService(attachmentRepo, *dependencies)
	troubleTicketService := services.NewTroubleTicketService(troubleTicketRepo)

	// controllers
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	attachmentController := controllers.NewAttachmentController(attachmentService)
	troubleTicketController := controllers.NewTroubleTicketController(troubleTicketService)

	// routers
	appRouter := routers.NewAppRouter(appController)
	authRouter := routers.NewAuthRouter(authController, *dependencies)
	attachmentRouter := routers.NewAttachmentRouter(attachmentController, *dependencies)
	troubleTicketRouter := routers.NewTroubleTicketRouter(troubleTicketController, *dependencies)
	// main router (putting all together)
	mainRouter := routers.NewMainRouter(
		*dependencies,
		appRouter,
		authRouter,
		troubleTicketRouter,
		attachmentRouter,
	)

	// start server
	if err := mainRouter.StartServer(*dependencies); err != nil {
		log.Panic("Error starting server:", err)
	}
}
