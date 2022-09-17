package settings

import (
	"bytes"
	"testing"
)

func TestDecodeJsonConfig(t *testing.T) {

	wants := Configuration{
		Pwm_pin: 18,
		Temperatures: [4]uint32{
			70, 65, 60, 55,
		},
	}

	var buffer bytes.Buffer
	buffer.WriteString("{\"pwm_pin\": 18, \"temperatures\": [ 70, 65, 60, 55]}")
	got, err := decodeJsonConfig(&buffer)
	if err != nil {
		t.Errorf("Failed to read json: %v\n", err)
	}
	if got != wants {
		t.Errorf("wanted: %v got %v\n", wants, got)
	}
}