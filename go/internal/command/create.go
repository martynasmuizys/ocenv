package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/martynasmuizys/ocenv/internal/log"
	"github.com/martynasmuizys/ocenv/internal/util"
)

func Create(name string) {
	r := bufio.NewReader(os.Stdin)

	if len(name) == 0 {
		log.Printf("Enter cluster name: ")
		input, err := r.ReadString('\n')

		if err != nil {
			log.Fatal(fmt.Errorf("Failed to read user input"))
		}

		name = input[:len(input)-1]
	}

	cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), name)
	cmd := Command{kubeConfig: cfg}

	if _, err := os.Stat(cfg); err == nil {
		if err := cmd.TermNoBuffering(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Configuration for cluster '%s' already exists! Overwrite? [y/N] ", name)

		char, _ := r.ReadByte()

		fmt.Println()
		if err := cmd.TermRestore(); err != nil {
			log.Fatal(err)
		}
		if char != 'y' && char != 'Y' {
			os.Exit(0)
		}
	}
	log.Printf("Enter cluster URL: ")
	input, err := r.ReadString('\n')
	url := strings.TrimSpace(input[:len(input)-1])

	if err != nil {
		log.Fatal(fmt.Errorf("Failed to read user input"))
	}

	log.Println("Trying to authenticate via browser...")
	if err := cmd.Oc(Login, url); err != nil {
		log.Fatal(err)
	}
	// 1 day
	expirationTimestamp := time.Now().Unix() + 86400
	kubeCfg, err := util.ParseConfig(cfg)

	if err != nil {
		log.Fatal(err)
	}

	kubeCfg.OcenvTokenExpires = expirationTimestamp
	err = util.SaveConfig(kubeCfg, cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cluster configuration created at '%s'\n", cfg)

}
