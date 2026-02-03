package util

import "fmt"

type CommandError struct {
	Cmd string
	Msg string
}

func (e *CommandError) Error() string {
	return fmt.Sprintf("%s - %s", e.Cmd, e.Msg)
}
