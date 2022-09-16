package settings

import (
	"encoding/json"
	"os"
	"log"
)

func JsonConfig(configLoc string) Configuration {

	logger := log.Default()

	file, err := os.Open(configLoc)
	if err != nil {
		log.Fatalf("failed to load %v: %v", configLoc, err)
	}
	defer file.Close()

	logger.Printf("Attempting to read config file: %v", configLoc)

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode %v: %v", configLoc, err)
	}

	logger.Printf("Successfully loaded config %v as: %v\n", configLoc, config)

	return config
}

type Configuration struct {
	Pwm_pin	uint32
	Temperatures [4]uint32
}
