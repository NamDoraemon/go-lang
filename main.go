package main

import (
	"fm.auth/cmd"
	"fmt"
)

func main() {
	fmt.Println("$ > STARTED")
	err := cmd.SetupCmd()
	if err != nil {
		panic(err)
	}
}
