package fan

import (
	"bytes"
	"testing"
)

func TestReadSysTempFile(t *testing.T) {

	var buffer bytes.Buffer
	buffer.WriteString("54043\n")
	var wants float64 = 54043

	got, err := readSysTempFile(&buffer)
	if err != nil {
		t.Fatal(err)
	}

	if got != wants {
		t.Fatalf("wanted: %v got: %v", wants, got)
	}
}