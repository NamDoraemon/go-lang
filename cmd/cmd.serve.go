package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namth/go-examples/config"
	"github.com/namth/go-examples/routes"
	"github.com/urfave/cli/v2"
)

func HTTPServe(ctx *cli.Context) error {
	configs := config.GetConfig()
	fmt.Println(configs)
	app := gin.New()
	app.RedirectTrailingSlash = false
	app.Use(gin.Recovery(), gin.Logger())
	fmt.Println(configs.HttpPort)
	routes.InitRoutePing(app)
	routes.InitRouteAuth(app)
	return app.Run(":" + configs.HttpPort)
}
