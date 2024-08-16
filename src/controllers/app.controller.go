package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AppController interface defines the contract for application controller.
type AppController interface {
	Index(context *gin.Context)
}

type appController struct{}

// Index handles the root URL of the application and returns a simple JSON response.
func (appCtl *appController) Index(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"message": "Trouble Ticket API"},
	)
}

// NewAppController creates a new instance of the appController struct & returns it as an AppController interface.
func NewAppController() AppController {
	return &appController{}
}
