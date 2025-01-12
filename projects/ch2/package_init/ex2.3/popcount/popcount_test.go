package popcount_test

import (
	"main/popcount"
	"testing"
)

const k uint64 = 1024

func TestFunction15(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(15))  // Output 4 (binary: 0x1111)
	result := popcount.PopCount(15)
	if result != 4 {
		t.Errorf("PopCount(15) = %d; want 4", result)
	}
}

func TestFunction1k(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(1024))        // Output 1 (binary: 0x1000 0000)
	result := popcount.PopCount(1 * k)
	if result != 1 {
		t.Errorf("PopCount(1k) = %d; want 1", result)
	}
}

func TestFunction256(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(256)) // Output 1 (binary: 0x0001 0000 0000)
	result := popcount.PopCount(256)
	if result != 1 {
		t.Errorf("PopCount(256) = %d; want 1", result)
	}
}

func TestFunction4k(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(4096))        // Output 1 (binary: 0x1000 0000)
	result := popcount.PopCount(4 * k)
	if result != 1 {
		t.Errorf("PopCount(4k) = %d; want 1", result)
	}
}

func TestFunction4km1(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(4095))        // Output 1 (binary: 0x0FFF FFFF)
	result := popcount.PopCount(4*k - 1)
	if result != 12 {
		t.Errorf("PopCount(4*k-1) = %d; want 12", result)
	}
}

func TestFunction64k(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(64 * 1024))   // Output 1 (binary: 0x0001 0000)
	result := popcount.PopCount(64 * k)
	if result != 1 {
		t.Errorf("PopCount(64k) = %d; want 1", result)
	}
}

func TestFunction64km1(t *testing.T) {
	// Test Logic
	// fmt.Println(popcount.PopCount(64*1024 - 1)) // Output 16 (binary: 0xFFFF)
	result := popcount.PopCount(64*k - 1)
	if result != 16 {
		t.Errorf("PopCount(64k-1) = %d; want 16", result)
	}
}

// demo a table test of values
var flagtests = []struct {
	name string
	in   uint64
	out  int
}{
	{"one", 1, 1},
	{"two", 2, 1},
	{"three", 3, 2},
	{"four", 4, 1},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			s := popcount.PopCount(tt.in)
			if s != tt.out {
				t.Errorf("Test name: %q, got %q, want %q", tt.name, s, tt.out)
			}
		})
	}
}
