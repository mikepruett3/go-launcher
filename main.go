package main

import (
	"log"
	"os/exec"

	"github.com/jinzhu/configor"
)

const (
	configFile = "config.yml"
)

var Config = struct {
	Browser struct {
		Exec  string
		Path  string
		Links []string
	}

	Programs []struct {
		Name string
		Exec string
		Path string
		Arg  string
	}
}{}

func main() {
	err := configor.Load(&Config, configFile)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(Config.Browser.Links); i++ {
		browser := exec.Command(Config.Browser.Exec, Config.Browser.Links[i])
		browser.Dir = Config.Browser.Path
		browserErr := browser.Start()
		if browserErr != nil {
			log.Fatal(browserErr)
		}
	}

	for i := 0; i < len(Config.Programs); i++ {
		cmd := exec.Command(Config.Programs[i].Exec, Config.Programs[i].Arg)
		cmd.Dir = Config.Programs[i].Path
		execErr := cmd.Start()
		if execErr != nil {
			log.Fatal(execErr)
		}
	}
}