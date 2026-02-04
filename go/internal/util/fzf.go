package util

import (
	"fmt"
	"os"
	"strings"

	fzf "github.com/junegunn/fzf/src"
	"github.com/martynasmuizys/ocenv/internal/log"
)

func Run(input []os.DirEntry, outputChan chan string) {
	inputChan := make(chan string)
	go func(entries *[]os.DirEntry) {
		for _, e := range *entries {
			inputChan <- strings.Split(e.Name(), ".")[0]
		}
		close(inputChan)
	}(&input)

	// Build fzf.Options
	options, err := fzf.ParseOptions(
		true, // whether to load defaults ($FZF_DEFAULT_OPTS_FILE and $FZF_DEFAULT_OPTS)
		[]string{"--reverse", "--border", "--height=40%"},
		// []string{"--border", "--height=40%"},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to parse `fzf` options: %v", err))
	}

	// Set up input and output channels
	options.Input = inputChan
	options.Output = outputChan

	// Run fzf
	_, err = fzf.Run(options)

	if err != nil {
		log.Fatal(fmt.Errorf("Failed to run `fzf`: %v", err))
	}
}
