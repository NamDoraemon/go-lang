package routes

import (
	"fm.auth/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouteAuth(route *gin.Engine) {
	authHandler := handlers.NewHandlerAuth()
	groupRoute := route.Group("/login")

	groupRoute.POST("", authHandler.LoginHandler)
}
