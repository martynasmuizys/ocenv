package command

import (
	"os"
	"os/exec"
	"time"

	"github.com/martynasmuizys/ocenv/internal/util"
)

type OcCommand int

const (
	Login OcCommand = iota
	Whoami
)

type TmuxCommand int

const (
	NewSession TmuxCommand = iota
	Switch
	ListSessions
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

func (c *Command) Tmux(tc TmuxCommand, args ...string) ([]byte, error) {
	var cmd *exec.Cmd
	name := "_ocenv_" + args[0]

	switch tc {
	case NewSession:
		cmd = exec.Command(
			"tmux", "new-session", "-ds", name, "-e", "KUBECONFIG="+c.kubeConfig)
	case Switch:
		if len(os.Getenv("TMUX")) == 0 {
			cmd = exec.Command("tmux", "attach-session", "-t", name)
		} else {
			cmd = exec.Command("tmux", "switch-client", "-t", name)
		}
	case ListSessions:
		cmd = exec.Command("tmux", "list-sessions")
	}

	out, err := cmd.Output()

	if err != nil {
		return nil, &util.CommandError{Cmd: "tmux", Msg: err.Error()}
	}

	return out, nil
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

func (c *Command) OcCheckToken(kubeCfg *util.KubeConfig) bool {
	err := c.Oc(Whoami, c.kubeConfig)
	if err != nil {
		// not an error. just checking whether token is valid
		return false
	}

	tokenExpires := kubeCfg.OcenvTokenExpires

	if ((int64(tokenExpires) - time.Now().Unix()) / 3600) < 8 {
		// 8 hour before expiration seems ok
		return false
	}

	return true
}
