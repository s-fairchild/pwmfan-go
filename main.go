package main

import (
	"log"
	"time"

	"github.com/s-fairchild/pwmfan-go/fan"
	"github.com/s-fairchild/pwmfan-go/settings"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

const cycleLength = 30000
const pwmClockFreq = 4 * cycleLength

func main() {

	config := settings.JsonConfig()

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
		runFan, dutyLength := fan.MonitorCpuTemp(config)
		if runFan {
			pin.DutyCycleWithPwmMode(dutyLength, 4, true)
		} else {
			pin.DutyCycle(0, 0)
		}
		// sleep for longest amount of time (4 minutes) when temperature is highest
		time.Sleep(time.Duration(dutyLength) * time.Minute)
	}
}
