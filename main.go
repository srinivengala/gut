package main

import (
	"fmt"
	"os"

	"github.com/srinivengala/gut/command"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Oops error : ", err)
			os.Exit(1)
		}
	}()

	os.Exit(command.Run(os.Args[1:]))
}
