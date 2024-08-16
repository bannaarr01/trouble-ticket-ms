package main

import (
	"log"
	"trouble-ticket-ms/src/controllers"
	"trouble-ticket-ms/src/routers"
)

func bootstrap() {
	// controllers
	appController := controllers.NewAppController()

	// routers
	appRouter := routers.NewAppRouter(appController)
	mainRouter := routers.NewMainRouter(appRouter)

	// start server
	if err := mainRouter.StartServer(); err != nil {
		log.Panic("Error starting server:", err)
	}
}
