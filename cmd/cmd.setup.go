package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
	"os"
)

func SetupCmd() error {
	cmdApp := cli.NewApp()
	cmdApp.EnableBashCompletion = true
	cmdApp.Action = func(c *cli.Context) error {
		return errors.New("please enter your command")
	}
	cmdApp.Commands = []*cli.Command{
		{Name: "serve", Usage: "Start up running app by serve", Action: HTTPServe},
		{Name: "grpc", Usage: "Start up running app by grpc", Action: GRPCServe},
		{Name: "run", Usage: "Start up running app by serve and grpc", Action: AppServe},
	}
	return cmdApp.Run(os.Args)
}
