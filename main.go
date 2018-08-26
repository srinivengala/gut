package main

import (
	"os"

	"github.com/srinivengala/gut/command"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
