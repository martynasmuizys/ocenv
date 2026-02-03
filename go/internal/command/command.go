package command

import (
	"os/exec"

	"github.com/martynasmuizys/ocenv/internal/util"
)

type OcCommand int

const (
	Login OcCommand = iota
	Whoami
)

type Command struct {
	ttyState   string
	kubeConfig string
}

func (c *Command) TermNoBuffering() error {
	// idk
	if len(c.ttyState) == 0 {
		out, err := exec.Command("stty", "-f", "/dev/tty", "-g").Output()
		c.ttyState = string(out)
		if err != nil {
			return &util.CommandError{Cmd: "stty", Msg: "Could not get default state"}
		}
	}

	err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		return &util.CommandError{Cmd: "stty", Msg: "Could not disable buffering"}
	}

	return nil
}

func (c *Command) TermRestore() error {
	err := exec.Command("stty", "-f", "/dev/tty", c.ttyState).Run()
	if err != nil {
		return &util.CommandError{Cmd: "stty", Msg: "Could not restore default state"}
	}

	return nil
}

func (c *Command) Tmux(args ...string) error {
	err := exec.Command("tmux", args...).Run()
	if err != nil {
		return &util.CommandError{Cmd: "tmux", Msg: err.Error()}
	}

	return nil
}

func (c *Command) Oc(occ OcCommand, args ...string) error {
	var cmd *exec.Cmd

	switch occ {
	case Login:
		cmd = exec.Command("oc", "login", "-w", "--server", args[0])
	case Whoami:
		cmd = exec.Command("oc", "whoami")
	}
	cmd.Args = append(cmd.Args, "--kubeconfig", c.kubeConfig)

	err := cmd.Run()

	if err != nil {
		return &util.CommandError{Cmd: "oc", Msg: err.Error()}
	}

	return nil
}
