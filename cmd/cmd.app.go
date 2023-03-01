package cmd

import "github.com/urfave/cli/v2"

func AppServe(ctx *cli.Context) error {
	go func() {
		err := GRPCServe(ctx)
		if err != nil {
			return
		}
	}()

	return HTTPServe(ctx)
}
