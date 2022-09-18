package fan

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/s-fairchild/pwmfan-go/settings"
)

const sysTemperatureFile = "/sys/class/thermal/thermal_zone0/temp"

// MonitorCpuTemp calls readCpuTemp and compares the temperature to the configuration thresholds
// returning a bool to turn the fan on/off, and the duty length when true
// When false returns dutylength of 5 for main loop to calculate sleep time
// float64 returned is the cpu temperature
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
func MonitorCpuTemp(c *settings.Configuration) (bool, uint32, float64) {

	const defaultSleepTime = 2

	if len(c.Temperatures) < 1 {
		fmt.Printf("temperatures configuration section must contain 4 values\n Example: \"temperatures\": [ 50, 55, 60, 65 ]\n")
		os.Exit(1)
	}

	file, err := os.Open(sysTemperatureFile)
	if err != nil {
		fmt.Printf("Unable to read %v: %v", sysTemperatureFile, err)
		os.Exit(1)
	}

	temp, err := readSysTempFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cTemp := temp / 1000
	
	// Requires c.Temperatures to be in descending order
	for dutyLength, threshold := range c.Temperatures {
		if uint32(cTemp) >= threshold {
			finalDutyLength := calculateDutyLengthAbs(dutyLength, len(c.Temperatures))
			return true, finalDutyLength, cTemp
		}
	}

	// Prevent endless loop that sleeps for 0 minutes by returning 2
	return false, defaultSleepTime, cTemp
}

// calculateDutyLengthAbs calculates the reversed duty length
func calculateDutyLengthAbs(dutyLength int, configLength int) uint32 {

	dutyLengthAbs := math.Abs(float64((dutyLength - configLength)))
	if math.IsNaN(dutyLengthAbs) {
		fmt.Printf("Failed to calculate dutylength absolute value: %v - %v", dutyLength, configLength)
		os.Exit(1)
	}

	return uint32(dutyLengthAbs)
}

// readSysTempFile parses the kernel system temperature
//
// file is modified to remove new line character and converted to a float
func readSysTempFile(reader io.Reader) (float64, error) {

	contents, err := io.ReadAll(reader)
	if err != nil {
		return 0, err
	}

	// remove line break from string
	strTempInt := string(contents[:len(contents)-1])
	sysTemp, err := strconv.ParseFloat(strTempInt, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %v into integer: %v", sysTemp, err)
	}
	
	return sysTemp, nil
}

