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
	// repositories
	troubleTicketRepo := repositories.NewTroubleTicketRepository(dbConn.DB)

	// services
	troubleTicketService := services.NewTroubleTicketService(troubleTicketRepo)

	// controllers
	appController := controllers.NewAppController()
	troubleTicketController := controllers.NewTroubleTicketController(troubleTicketService)

	// routers
	appRouter := routers.NewAppRouter(appController)
	troubleTicketRouter := routers.NewTroubleTicketRouter(troubleTicketController)
	mainRouter := routers.NewMainRouter(appRouter, troubleTicketRouter)

	// start server
	if err := mainRouter.StartServer(); err != nil {
		log.Panic("Error starting server:", err)
	}
}
