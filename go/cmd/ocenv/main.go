package main

import (
	"flag"
	"fmt"

	"github.com/martynasmuizys/ocenv/internal/command"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "":
		break
	case "create":
		command.Create(flag.Arg(1))
	case "help":
		command.Help()
	default:
		command.Help()
		fmt.Printf("error: unknown command '%s'\n", flag.Arg(0))
	}
}
