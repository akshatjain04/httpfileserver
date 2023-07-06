package httpfileserver

import (
	"testing"
)

type fileServer struct {
	optionMaxBytesPerFile int64
}

func OptionMaxBytes(maxBytes int64) func(*fileServer) {
	return func(fs *fileServer) {
		fs.optionMaxBytesPerFile = maxBytes
	}
}

// TestOptionMaxBytes_1b8c6fdb68 tests the OptionMaxBytes function.
func TestOptionMaxBytes_1b8c6fdb68(t *testing.T) {
	// Test case 1: Normal case
	fs := &fileServer{}
	option := OptionMaxBytes(1000)
	option(fs)
	if fs.optionMaxBytesPerFile != 1000 {
		t.Errorf("Expected optionMaxBytesPerFile to be %d, but got %d", 1000, fs.optionMaxBytesPerFile)
	}

	// Test case 2: Setting optionMaxBytesPerFile to zero
	fs = &fileServer{}
	option = OptionMaxBytes(0)
	option(fs)
	if fs.optionMaxBytesPerFile != 0 {
		t.Errorf("Expected optionMaxBytesPerFile to be %d, but got %d", 0, fs.optionMaxBytesPerFile)
	}

	// Test case 3: Setting optionMaxBytesPerFile to negative value
	fs = &fileServer{}
	option = OptionMaxBytes(-1000)
	option(fs)
	if fs.optionMaxBytesPerFile != -1000 {
		t.Errorf("Expected optionMaxBytesPerFile to be %d, but got %d", -1000, fs.optionMaxBytesPerFile)
	}
}
