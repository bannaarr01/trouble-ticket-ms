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
	troubleTicketRepo := repositories.NewTroubleTicketRepository(dbConn.DB)

	// services
	authService := services.NewAuthService(*dependencies)
	troubleTicketService := services.NewTroubleTicketService(troubleTicketRepo)

	// controllers
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	troubleTicketController := controllers.NewTroubleTicketController(troubleTicketService)

	// routers
	appRouter := routers.NewAppRouter(appController)
	authRouter := routers.NewAuthRouter(authController, *dependencies)
	troubleTicketRouter := routers.NewTroubleTicketRouter(troubleTicketController)
	// main router (putting all together)
	mainRouter := routers.NewMainRouter(*dependencies, appRouter, authRouter, troubleTicketRouter)

	// start server
	if err := mainRouter.StartServer(*dependencies); err != nil {
		log.Panic("Error starting server:", err)
	}
}
