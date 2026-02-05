package command

import (
	"fmt"
	"os"
	"sync"

	"github.com/martynasmuizys/ocenv/internal/log"
	"github.com/martynasmuizys/ocenv/internal/util"
)

func Use(name string) {
	var wg sync.WaitGroup
	if len(name) == 0 {
		outputChan := make(chan string)
		wg.Go(func() {
			s := <-outputChan
			cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), s)
			err := startSession(cfg, s)
			if err != nil {
				log.Fatal(err)
			}
		})

		dirs, err := os.ReadDir(fmt.Sprintf("%s/.kube/ocenv", os.Getenv("HOME")))
		if err != nil {
			log.Fatal(fmt.Errorf("Failed to read environment directory"))
		}
		code, err := util.GetEnvironments(dirs, outputChan)
		// sigint
		if (code == 130) && (err == nil) {
			wg.Done()
		}
	} else {
		cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), name)
		startSession(cfg, name)
	}
	wg.Wait()
}

func startSession(cfg string, name string) error {
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
		return fmt.Errorf("Failed to parse config file: %v", err)
	}
	valid := cmd.OcCheckToken(kubeCfg)
	if !valid {
		err := cmd.Oc(Login, kubeCfg.Clusters[0].Cluster.Server)
		if err != nil {
			return fmt.Errorf("Failed to login: %v", err)
		}
	}

	if !exists {
		if _, err := cmd.Tmux(NewSession, name); err != nil {
			return fmt.Errorf("Failed to create tmux session: %v", err)
		}
	}

	if _, err := cmd.Tmux(Switch, name); err != nil {
		return fmt.Errorf("Failed to switch to tmux session: %v", err)
	}

	return nil
}
