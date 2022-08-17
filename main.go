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
		Profile string
		Links []string
	}

	Programs []struct {
		Name     string
		Exec     string
		StartDir string
		Arg      string
	}

	ClickOnce []struct {
		Name     string
		Exec     string
		StartDir string
		Arg      string
	}
}{}

func main() {
	err := configor.Load(&Config, configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Start up Chrome Session using Specified Profile first
	browser := exec.Command(Config.Browser.Exec, "--profile-directory=", Config.Browser.Profile)
	browserErr := browser.Start()
	if browserErr != nil {
		log.Fatal(browserErr)
	}

	// Then load in each url
	for i := 0; i < len(Config.Browser.Links); i++ {
		browser := exec.Command(Config.Browser.Exec, "-url", Config.Browser.Links[i])
		browserErr := browser.Start()
		if browserErr != nil {
			log.Fatal(browserErr)
		}
	}

	for i := 0; i < len(Config.Programs); i++ {
		cmd := exec.Command(Config.Programs[i].Exec, Config.Programs[i].Arg)
		cmd.Dir = Config.Programs[i].StartDir
		execErr := cmd.Start()
		if execErr != nil {
			log.Fatal(execErr)
		}
	}

	//for i := 0; i < len(Config.ClickOnce); i++ {
	//	co := exec.Command("rundll32.exe", "dfshim.dll", ",", "ShOpenVerbApplication", Config.ClickOnce[i].Exec, Config.ClickOnce[i].Arg)
	//	co.Dir = Config.ClickOnce[i].StartDir
	//	coErr := co.Start()
	//	if coErr != nil {
	//		log.Fatal(coErr)
	//	}
	//}
}
