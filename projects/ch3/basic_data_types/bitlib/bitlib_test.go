package bitlib

import (
	"testing"
)

func TestAnd(t *testing.T) {
	result := BitAnd(0xFF, 0x03)
	if result != 0x03 {
		t.Error("BIT AND fail")
	}
}

func TestClear(t *testing.T) {
	result := BitClear(0xFF, 0x03)
	if result != 0xFC {
		t.Error("BIT CLEAR fail")
	}
}
