package command

import (
	"fmt"
)

func Help() {
	fmt.Printf("Usage: ocenv [SUBCOMMAND] [OPTIONS]\n")
	fmt.Printf("Subcommands:\n")
	fmt.Printf("  %-14s  %s\n", "create <NAME>", "Creates a new environment for a cluster")
	fmt.Printf("  %-14s  %s\n", "info", "Prints information about environment")
	fmt.Printf("  %-14s  %s\n", "list", "Lists all active sessions for configured environments")
	fmt.Printf("  %-14s  %s\n", "rm <NAME>", "Removes environment")
	fmt.Printf("  %-14s  %s\n", "use <NAME>", "Initializes a TMUX session with selected environment")
	fmt.Printf("Options:\n")
	fmt.Printf("  %-14s  %s\n", "-h, --help", "Prints this help message")
	fmt.Printf("  %-14s  %s\n", "-v, --version", "Prints version information")
}
