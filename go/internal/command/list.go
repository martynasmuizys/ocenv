package command

import (
	"fmt"
	"os"

	"github.com/martynasmuizys/ocenv/internal/log"
)

func List(active bool) {
	if active {
		log.Println("active")
		return
	}

	dirs, err := os.ReadDir(fmt.Sprintf("%s/.kube/ocenv", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to read environment directory"))
	}

	log.Hprint("Environments")
	for _, d := range dirs {
		log.Println(d.Name())
	}

}
