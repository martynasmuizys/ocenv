package command

import (
	"fmt"
	"os"

	"github.com/martynasmuizys/ocenv/internal/log"
	"github.com/martynasmuizys/ocenv/internal/util"
)

func Use(name string) {
	if len(name) == 0 {
		outputChan := make(chan string)
		go func() {
			for s := range outputChan {
				cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), s)
				startSession(cfg, s)
			}
		}()

		dirs, err := os.ReadDir(fmt.Sprintf("%s/.kube/ocenv", os.Getenv("HOME")))
		if err != nil {
			log.Fatal(fmt.Errorf("Failed to read environment directory"))
		}
		util.Run(dirs, outputChan)
	} else {
		cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), name)
		startSession(cfg, name)
	}
}

func startSession(cfg string, name string) {
	cmd := Command{kubeConfig: cfg}
	var exists bool

	out, err := cmd.Tmux(ListSessions, name)
	if err != nil {
		log.Fatal(err)
		// log.Fatal(fmt.Errorf("Failed to create tmux session: %v", err))
	}

	for _, s := range util.ParseSessions(string(out)) {
		if s == "_ocenv_"+name {
			exists = true
		}
	}

	kubeCfg, err := util.ParseConfig(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to parse config file: %v", err))
	}
	valid := cmd.OcCheckToken(kubeCfg)
	if !valid {
		cmd.Oc(Login, kubeCfg.Clusters[0].Cluster.Server)
	}

	if !exists {
		if _, err := cmd.Tmux(NewSession, name); err != nil {
			log.Fatal(fmt.Errorf("Failed to create tmux session: %v", err))
		}
	}

	if _, err := cmd.Tmux(Switch, name); err != nil {
		log.Fatal(fmt.Errorf("Failed to switch to tmux session: %v", err))
	}
}
