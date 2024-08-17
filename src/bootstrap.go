package main

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"log"
	"trouble-ticket-ms/src/config"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/db"
	"trouble-ticket-ms/src/repositories"
	"trouble-ticket-ms/src/routers"
	"trouble-ticket-ms/src/services"
)

func bootstrap(dbConn *db.DB) {
	// config
	cfg := config.New()

	// keycloak
	client := gocloak.NewClient(cfg.KEYCLOAK.Host)
	ctx := context.Background()

	// repositories
	troubleTicketRepo := repositories.NewTroubleTicketRepository(dbConn.DB)

	// services
	authService := services.NewAuthService(client, ctx, cfg.KEYCLOAK)
	troubleTicketService := services.NewTroubleTicketService(troubleTicketRepo)

	// controllers
	appController := controllers.NewAppController()
	authController := controllers.NewAuthController(authService)
	troubleTicketController := controllers.NewTroubleTicketController(troubleTicketService)

	// routers
	appRouter := routers.NewAppRouter(appController)
	authRouter := routers.NewAuthRouter(authController)
	troubleTicketRouter := routers.NewTroubleTicketRouter(troubleTicketController)
	// main router (putting all together)
	mainRouter := routers.NewMainRouter(appRouter, authRouter, troubleTicketRouter)

	// start server
	if err := mainRouter.StartServer(); err != nil {
		log.Panic("Error starting server:", err)
	}
}
