package settings

import (
	"encoding/json"
	"os"
	"log"
)

const configFile = "pwm-conf.json"

func JsonConfig() Configuration {

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("failed to load %v: %v", configFile, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode %v: %v", configFile, err)
	}

	return config
}

type Configuration struct {
	Pwm_pin	uint32
}
