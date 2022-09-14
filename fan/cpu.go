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
func MonitorCpuTemp(c settings.Configuration) (bool, uint32) {

	if len(c.Temperatures) < 1 {
		log.Fatalln("temperatures configuration section must contain 4 values\n Example: \"temperatures\": [ 50, 55, 60, 65 ]")
	}

	cTemp, err := readCpuTemp()
	fmt.Printf("Current CPU temperature is %.f\n", cTemp)
	if err != nil {
		log.Fatalln(err)
	}

	// Requires c.Temperatures to be in descending order
	for dutyLength, threshold := range c.Temperatures {
		fmt.Printf("Comparing cpu temperature %v to %v, index: %v\n", uint32(cTemp), threshold, dutyLength)
		if uint32(cTemp) > threshold {
			fmt.Printf("%v is greater than threshold %v\n", uint32(cTemp), threshold)
			finalDutyLength := calculateDutyLengthAbs(dutyLength, len(c.Temperatures))

			return true, finalDutyLength
		}
	}

	return false, 0
}

// calculateDutyLengthAbs calculates the reversed duty length
func calculateDutyLengthAbs(dutyLength int, configLength int) uint32 {

	dutyLengthAbs := math.Abs(float64((dutyLength - configLength)))
	if math.IsNaN(dutyLengthAbs) {
		log.Fatalf("Failed to calculate dutylength absolute value: %v - %v", dutyLength, configLength)
	}

	fmt.Printf("returning dutylength of: %v\n", dutyLengthAbs)
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
