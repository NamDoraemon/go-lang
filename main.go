package main

import (
	"github.com/namth/go-examples/cmd"
)

func main() {
	err := cmd.SetupCmd()
	if err != nil {
		panic(err)
	}
}
