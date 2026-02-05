package util

import (
	"fmt"
	"os"
	"strings"

	fzf "github.com/junegunn/fzf/src"
)

func GetEnvironments(input []os.DirEntry, outputChan chan string) (int, error) {
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
		return 1, fmt.Errorf("Failed to parse `fzf` options: %v", err)
	}

	// Set up input and output channels
	options.Input = inputChan
	options.Output = outputChan

	// Run fzf
	code, err := fzf.Run(options)

	if err != nil {
		return 1, fmt.Errorf("Failed to run `fzf`: %v", err)
	}
	return code, err
}
