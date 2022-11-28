package main

import (
	"os"

	"com.son.server.base/cmd/baseserver/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
