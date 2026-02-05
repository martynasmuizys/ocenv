package command

import (
	"fmt"
	"os"

	"github.com/martynasmuizys/ocenv/internal/log"
)

func Rm(name string) {
	if len(name) == 0 {
		log.Fatal(fmt.Errorf("No environment name provided"))
	}

	cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), name)

	if _, err := os.Stat(cfg); err == nil {
		if err := os.Remove(cfg); err != nil {
			log.Fatal(fmt.Errorf("Failed to delete environment: %v", err))
		}

		log.Printf("Environment '%s' was removed\n", name)
	} else {
		log.Fatal(fmt.Errorf("Environment does not exist"))
	}
}
