package fan

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"

	"github.com/s-fairchild/pwmfan-go/settings"
)

// MonitorCpuTemp calls readCpuTemp and compares the temperature to the configuration thresholds
// returning a bool to turn the fan on/off, and the duty length when true
// When falase returns dutylength of 5 for main loop to calculate sleep time
//
// dutylength is a fraction of 4
//
// dutylength 1 = 1/4 power or 25%
//
// dutylength 2 = 2/4 power or 50%
//
// dutylength 3 = 3/4 power or 75%
//
// dutylength 4 = 4/4 power or 100%
func MonitorCpuTemp(c settings.Configuration) (bool, uint32) {

	const defaultSleepTime = 2

	if len(c.Temperatures) < 1 {
		log.Fatalln("temperatures configuration section must contain 4 values\n Example: \"temperatures\": [ 50, 55, 60, 65 ]")
	}

	cTemp, err := readCpuTemp()
	fmt.Printf("CPU temperature is %.2f\n", cTemp)
	if err != nil {
		log.Fatalln(err)
	}

	// Requires c.Temperatures to be in descending order
	for dutyLength, threshold := range c.Temperatures {
		if uint32(cTemp) >= threshold {
			finalDutyLength := calculateDutyLengthAbs(dutyLength, len(c.Temperatures))
			fmt.Printf("Running fan at %v/4 power\n", finalDutyLength)

			return true, finalDutyLength
		}
	}

	// Prevent endless loop that sleeps for 0 minutes by returning 2
	return false, defaultSleepTime
}

// calculateDutyLengthAbs calculates the reversed duty length
func calculateDutyLengthAbs(dutyLength int, configLength int) uint32 {

	dutyLengthAbs := math.Abs(float64((dutyLength - configLength)))
	if math.IsNaN(dutyLengthAbs) {
		log.Fatalf("Failed to calculate dutylength absolute value: %v - %v", dutyLength, configLength)
	}

	return uint32(dutyLengthAbs)
}

// readCpuTemp extracts and converts output from vcgencmd
// extracted float is rounded to nearest value
func readCpuTemp() (float64, error) {

	cmd := exec.Command("vcgencmd", "measure_temp")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return 0, fmt.Errorf("error running vcgencmd: %v", err)
	}

	output := out.String()

	// Extract temperature substring
	// example of output that we are slicing for the substring: temp=61.8'C
	outputSubStr := output[5:len(output)-3]
	cpuTemp, err := strconv.ParseFloat(outputSubStr, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse vcgencmd output: %v", err)
	}

	return cpuTemp, nil
}
