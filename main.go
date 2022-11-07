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
		Exec     string
		Profile  string
		StartDir string
		Links    []string
	}

	Programs []struct {
		Exec     string
		StartDir string
		Arg      string
	}

	Elevated []struct {
		Exec     string
		StartDir string
		Arg      string
	}

	Pwa []struct {
		Exec     string
		StartDir string
		Profile  string
		Appid    string
	}

	//ClickOnce []struct {
	//	Exec     string
	//	StartDir string
	//	Arg      string
	//}
}{}

func main() {
	err := configor.Load(&Config, configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Browser
	for i := 0; i < len(Config.Browser.Links); i++ {
		browser := exec.Command(Config.Browser.Exec, "--profile-directory="+Config.Browser.Profile, "-url", Config.Browser.Links[i])
		browserErr := browser.Start()
		if browserErr != nil {
			log.Fatal(browserErr)
		}
	}

	// Programs
	for i := 0; i < len(Config.Programs); i++ {
		cmd := exec.Command(Config.Programs[i].Exec, Config.Programs[i].Arg)
		cmd.Dir = Config.Programs[i].StartDir
		execErr := cmd.Start()
		if execErr != nil {
			log.Fatal(execErr)
		}
	}

	// Elevated Programs
	for i := 0; i < len(Config.Elevated); i++ {
		cmd := exec.Command("sudo", Config.Elevated[i].Exec, Config.Elevated[i].Arg)
		cmd.Dir = Config.Elevated[i].StartDir
		execErr := cmd.Start()
		if execErr != nil {
			log.Fatal(execErr)
		}
	}

	// PWA Apps
	for i := 0; i < len(Config.Pwa); i++ {
		cmd := exec.Command(Config.Pwa[i].Exec, "--profile-directory="+Config.Pwa[i].Profile, "--app-id="+Config.Pwa[i].Appid)
		cmd.Dir = Config.Pwa[i].StartDir
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
