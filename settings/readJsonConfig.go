package settings

import (
	"encoding/json"
	"os"
	"fmt"
	"io"
)

func ReadFanConfig(configLoc string) Configuration {

	fmt.Printf("Attempting to read config file: %v\n", configLoc)
	file, err := os.Open(configLoc)
	if err != nil {
		fmt.Printf("failed to load %v: %v\n", configLoc, err)
		os.Exit(1)
	}
	defer file.Close()

	config, err := decodeJsonConfig(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Successfully loaded config %v as: %v\n", configLoc, config)

	return config
}

func decodeJsonConfig(file io.Reader) (Configuration, error) {

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		return Configuration{}, fmt.Errorf("failed to decode file: %v", err)
	}

	return config, nil
}

type Configuration struct {
	Pwm_pin	uint32
	Temperatures [4]uint32
}
