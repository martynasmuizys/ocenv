package util

import (
	"strings"
)

func ParseSessions(sessions string) []string {
	sessionsArray := strings.Split(sessions, "\n")

	for i, s := range sessionsArray {
		sessionsArray[i] = strings.Split(s, ":")[0]
	}

	return sessionsArray
}
