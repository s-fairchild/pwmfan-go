package main

import (
	"fmt"
	"log"
	"time"

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

	cycleLengths := []uint32{1, 2, 3, 4}
	for _, cycleLength := range cycleLengths {
		fmt.Printf("Duty cycle=%v\n", cycleLength)
		pin.DutyCycleWithPwmMode(cycleLength, 4, true)
		time.Sleep(10 * time.Second)
	}
	pin.DutyCycle(0, 0)
}
