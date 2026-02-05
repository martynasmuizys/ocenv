package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/martynasmuizys/ocenv/internal/command"
)

func main() {
	defaultSet := flag.NewFlagSet("", flag.ExitOnError)
	defaultSet.Usage = func() {
		command.Help()
	}
	versionFlag := defaultSet.BoolP("version", "v", false, "Prints version information")

	// List subcommand
	listSet := flag.NewFlagSet("list", flag.ExitOnError)
	listSet.Usage = func() {
		command.ListHelp()
	}
	activeFlag := listSet.BoolP("active", "a", false, "Lists all active sessions")

	if len(os.Args) < 2 {
		command.Use("")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "":
		break
	case "create":
		var name string
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		command.Create(name)
	case "info":
		var name string
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		command.Info(name)
	case "list":
		if err := listSet.Parse(os.Args[2:]); err == nil {
			command.List(*activeFlag)
		}
	case "rm":
		var name string
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		command.Rm(name)
	case "use":
		var name string
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		command.Use(name)
	case "help":
		command.Help()
	case "version":
		command.Version()
	default:
		defaultSet.Parse(os.Args[1:])

		if *versionFlag {
			command.Version()
			os.Exit(0)
		}
		command.Help()
		fmt.Printf("error: unknown command '%s'\n", flag.Arg(0))
	}
}
