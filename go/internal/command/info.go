package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/martynasmuizys/ocenv/internal/log"
	"github.com/martynasmuizys/ocenv/internal/util"
)

func Info(name string) {
	if len(name) == 0 {
		log.Fatal(fmt.Errorf("No environment name provided"))
	}

	cfg := fmt.Sprintf("%s/.kube/ocenv/%s.yaml", os.Getenv("HOME"), name)

	kubeCfg, err := util.ParseConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	tokenExpires := (int64(kubeCfg.OcenvTokenExpires) - time.Now().Unix()) / 3600
	var hoursLeft string

	if tokenExpires <= 0 {
		hoursLeft = "expired"
	} else {
		hoursLeft = strconv.FormatInt(tokenExpires, 10)

		if hoursLeft == "1" {
			hoursLeft += " hour"
		} else {
			hoursLeft += " hours"
		}
	}

	log.Hprint(name)
	log.Printf("%-20s %s\n", "Cluster:", strings.Split(kubeCfg.CurrentContext, "/")[1])
	log.Printf("%-20s %s\n", "User:", strings.Split(kubeCfg.CurrentContext, "/")[2])
	log.Printf("%-20s %s\n", "Session expires:", hoursLeft)
	log.Printf("%-20s %s\n", "Namespace:", strings.Split(kubeCfg.CurrentContext, "/")[0])
}
