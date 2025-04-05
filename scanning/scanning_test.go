package scanning_test

import (
	"apothecary-journal/scanning"
	"testing"
)

func TestDecodeImage(t *testing.T) {
	expected := "png"
	_, actual, _ := scanning.DecodeImage("../images/training/abutilon.png")
	if expected != actual {
		t.Fatalf("Expected: %s | Actual: %s", expected, actual)
	}
}
