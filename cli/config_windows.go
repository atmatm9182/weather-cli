//go:build windows

package cli

import (
	"log"
	"os"
	"path"
)

const configDirName = "Weather CLI"

func configDir() string {
	appData, ok := os.LookupEnv("APPDATA")
	if !ok {
		log.Fatal("The %APPDATA% is unset")
	}

	return path.Join(appData, configDirName)
}
