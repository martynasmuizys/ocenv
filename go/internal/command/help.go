package command

import (
	"fmt"
	"strings"
)

func Version() {
	fmt.Println("ocenv(go) 0.1.0")
}

func Help() {
	var b strings.Builder
	b.WriteString("Usage:\n")
	b.WriteString("  ocenv [SUBCOMMAND] [OPTIONS]\n")
	b.WriteRune('\n')
	b.WriteString("Subcommands:\n")
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "create <NAME>", "Creates a new environment for a cluster"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "info <NAME>", "Prints information about environment"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "list", "Lists all active sessions for configured environments"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "rm <NAME>", "Removes environment"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "use <NAME>", "Initializes a TMUX session with selected environment"))
	b.WriteRune('\n')
	b.WriteString("Options:\n")
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "-h, --help", "Prints this help message"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "-v, --version", "Prints version information"))

	fmt.Print(b.String())
}

func ListHelp() {
	var b strings.Builder

	b.WriteString("Usage: ocenv [SUBCOMMAND] [OPTIONS]\n")
	b.WriteRune('\n')
	b.WriteString("Options:\n")
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "-a, --active", "Lists all active sessions"))
	b.WriteString(fmt.Sprintf("  %-14s  %s\n", "-h, --help", "Prints this help message"))

	fmt.Print(b.String())
}
