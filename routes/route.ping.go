package routes

import (
	handlers "auth/handlers/ping"
	"github.com/gin-gonic/gin"
)

func InitRoutePing(route *gin.Engine) {
	pingHandler := handlers.NewHandlerPing()
	groupRoute := route.Group("/ping")

	groupRoute.GET("", pingHandler.PingHandler)
}
