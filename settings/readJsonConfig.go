package settings

import (
	"encoding/json"
	"os"
	"log"
	"fmt"
)

func JsonConfig(configLoc string) Configuration {

	file, err := os.Open(configLoc)
	if err != nil {
		log.Fatalf("failed to load %v: %v\n", configLoc, err)
	}
	defer file.Close()

	fmt.Printf("Attempting to read config file: %v\n", configLoc)

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode %v: %v", configLoc, err)
	}

	fmt.Printf("Successfully loaded config %v as: %v\n", configLoc, config)

	return config
}

type Configuration struct {
	Pwm_pin	uint32
	Temperatures [4]uint32
}
