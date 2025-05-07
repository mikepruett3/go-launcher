//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/go-launcher.exe.manifest

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func checkDirExists(path string) bool {
	// Use os.Stat to get the file info
	_, err := os.Stat(path)

	// Check if the error is "not exist"
	if os.IsNotExist(err) {
		return false
	}

	// Check if it's actually a directory
	if err == nil {
		info, err := os.Stat(path)
		if err == nil && info.IsDir() {
			return true
		}
	}

	// Some other error (e.g., permission issue)
	return false
}

func checkFileExists(filename string) bool {
	// Use os.Stat to get the file info
	_, err := os.Stat(filename)

	// Check if it's actually a file
	if err == nil {
		return true // File exists
	}

	// Check if the error is "not exist"
	if os.IsNotExist(err) {
		return false // File does not exist
	}

	// Some other error (e.g., permission issue)
	return false
}

func main() {
	// Retrieve the ProgramFiles environment variable
	UserProfile := os.Getenv("UserProfile")
	if UserProfile == "" {
		fmt.Println("Error: USERPROFILE environment variable not found")
		return
	}

	viper.SetConfigName("config")                                 // name of config file (without extension)
	viper.SetConfigType("yaml")                                   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/go-launcher/")                      // path to look for the config file in
	viper.AddConfigPath(UserProfile + "\\.config\\go-launcher\\") // call multiple times to add many search paths
	viper.AddConfigPath("$HOME/.config/go-launcher/")             // ...
	viper.AddConfigPath("$ENV:LOCALAPPDATA\\go-launcher\\")       // ...
	viper.AddConfigPath("$HOME\\.config\\go-launcher\\")          // ...
	viper.AddConfigPath(".")                                      // optionally look for config in the working directory
	err := viper.ReadInConfig()                                   // Find and read the config file
	if err != nil {                                               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Browser Launch
	browsers := viper.GetStringMap("browser")
	if browsers == nil {
		log.Fatal("Browsers section not found in the config file!")
	}

	// Collect variables from config
	var browser string
	if viper.IsSet("browser.exec") {
		browser = viper.GetString("browser.exec")
	} else {
		log.Fatal("No executable for a browser is defined!")
	}

	// Collect args variable info from config, if exists
	var args string
	if viper.IsSet("browser.args") {
		args = viper.GetString("browser.args")
	}

	// Collect profile variable info from config, if exists
	var profile string
	if viper.IsSet("browser.profile") {
		profile = " --profile-directory=" + viper.GetString("browser.profile")
	}

	// Collect String Slice of Urls from config
	var links []string
	if viper.IsSet("browser.links") {
		links = viper.GetStringSlice("browser.links")
	} else {
		log.Fatal("No list of Urls found!")
	}

	switch true {
	case strings.Contains(browser, "chromium"):
		for _, link := range links {
			// Open Browser with specified URL
			cmd := exec.Command(browser, args, profile, link)

			// Detach from parent process
			//cmd.Stdout = os.Stdout
			//cmd.Stderr = os.Stderr

			// Start Program without waiting
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	case strings.Contains(browser, "firefox"):
		for _, link := range links {
			// Open Browser with specified URL
			cmd := exec.Command(browser, args, profile, " --url ", link)

			// Detach from parent process
			//cmd.Stdout = os.Stdout
			//cmd.Stderr = os.Stderr

			// Start Program without waiting
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	default:
		log.Fatal("Unknown Browser:")
	}

	// Program Launch
	programs := viper.GetStringMap("programs")
	if programs == nil {
		log.Fatal("Programs section not found in the config file!")
	}

	for program, details := range programs {
		var Exec string
		var Args string
		var StartDir string

		if Details, ok := details.(map[string]interface{}); ok {
			for key, value := range Details {
				switch key {
				case "exec":
					if exec, ok := value.(string); ok {
						if checkFileExists(exec) {
							Exec = exec
						} else {
							log.Fatal("File does not exist!")
						}
					}
				case "args":
					if args, ok := value.(string); ok {
						Args = args
					}
				case "start_dir":
					if start_dir, ok := value.(string); ok {
						if checkDirExists(start_dir) {
							StartDir = start_dir
						} else {
							log.Fatal("Directory does not exist!")
						}

					}
				default:
					log.Fatal("Unknown property:", key)
				}
			}
		} else {
			log.Fatal("Invalid structure for program:", program)
		}

		// Print Values from read config file
		//fmt.Println(Exec, Args, StartDir)

		// Open Program
		cmd := exec.Command(Exec, Args)

		// Detach from parent process
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr

		// Specify starting directory
		cmd.Dir = StartDir

		// Start Program without waiting
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
