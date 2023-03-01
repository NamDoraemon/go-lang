package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/namth/go-examples/handlers/auth"
)

func InitRouteAuth(route *gin.Engine) {
	authHandler := handlers.NewHandlerAuth()
	groupRoute := route.Group("/auth")

	groupRoute.POST("/login", authHandler.LoginHandler)
}
