//go:build linux

package cli

import (
	"fmt"
	"log"
	"os"
)

const configDirName = "weather-cli"

func configDir() string {
	home, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		return home
	}

	home, ok = os.LookupEnv("HOME")
	if !ok {
		log.Fatal("Neither $XDG_CONFIG_HOME or $HOME environment variables are not set, " +
			"please provide your api key explicitely using -key command line argument")
	}

	return fmt.Sprintf("%s/.config/%s", home, configDirName)
}
