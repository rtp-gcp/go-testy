package intlib

import (
	"testing"
)

func TestSomeFunc(t *testing.T) {
	result := SomeFunc()
	if result != 0x0 {
		t.Error("Test SomeFunc FAIL")
	}
}
