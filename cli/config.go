package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Config struct {
	ApiKey string
}

const configFileName = "config.json"

func parseSystemConfig() (config *Config, err error) {
	dir := configDir()

	configPath := path.Join(dir, configFileName)
	var f *os.File
	f, err = os.Open(configPath)
	if os.IsNotExist(err) {
		err = fmt.Errorf("Config file %s does not exist", configPath)
		return
	}
	if err != nil {
		return
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)

	if err != nil {
		config = nil
	}

	return
}
