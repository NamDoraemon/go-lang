package main

import (
	"fmt"
	"github.com/namth/go-examples/cmd"
)

func main() {
	fmt.Println("$ > STARTED")
	err := cmd.SetupCmd()
	if err != nil {
		panic(err)
	}
}
