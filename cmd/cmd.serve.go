package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/namth/go-examples/config"
	"github.com/namth/go-examples/routes"
	"github.com/urfave/cli/v2"
)

func HTTPServe(ctx *cli.Context) error {
	configs := config.GetConfigProduction()

	app := gin.New()
	app.RedirectTrailingSlash = false
	app.Use(gin.Recovery(), gin.Logger())

	routes.InitPingRoutes(app)
	return app.Run(":" + configs.HttpPort)
}
