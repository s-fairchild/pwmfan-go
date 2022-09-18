package main

import (
	"fmt"
	"log"
	"time"

	"github.com/s-fairchild/pwmfan-go/fan"
	"github.com/s-fairchild/pwmfan-go/settings"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

const (
	cycleLength      = 30000
	pwmClockFreq     = 4 * cycleLength
	configEnv        = "PWMFAN_CONFIG"
	defaultConfigLoc = "/usr/local/etc/pwmfan-conf.json"
)

func main() {

	configLoc := settings.GetConfigLocation(configEnv, defaultConfigLoc)
	config := settings.ReadFanConfig(configLoc)

	err := rpio.Open()
	if err != nil {
		log.Fatalf("Failed to open memory range in /dev/mem, %v", err)
	}
	defer rpio.Close()

	pin := rpio.Pin(config.Pwm_pin)
	pin.Mode(rpio.Pwm)
	pin.Pwm()
	pin.Freq(pwmClockFreq)
	rpio.StartPwm()

	for {
		runFan, dutyLength, cTemp := fan.MonitorCpuTemp(&config)
		fmt.Printf("CPU Temperature: %2.3f\n", cTemp)
		if runFan {
			pin.DutyCycleWithPwmMode(dutyLength, 4, true)
			fmt.Printf("Running fan at %v/4 power\n", dutyLength)
		} else {
			pin.DutyCycle(0, 0)
		}
		// sleep for 1-4 minutes before checking temperature again
		time.Sleep(time.Duration(dutyLength) * time.Minute)
	}
}
