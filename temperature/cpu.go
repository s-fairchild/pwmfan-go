package temperature

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

// ReadCpuTemp extracts and converts output from vcgencmd
//
// extracted float is rounded to nearest value
func ReadCpuTemp() (float64, error) {

	cmd := exec.Command("vcgencmd", "measure_temp")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return 0, fmt.Errorf("error running vcgencmd: %v", err)
	}

	output := out.String()

	// Extract temperature substring
	outputSubStr := output[5:len(output)-3]
	cpuTemp, err := strconv.ParseFloat(outputSubStr, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse vcgencmd output: %v", err)
	}

	return cpuTemp, nil
}
