package httpfileserver

import (
	"testing"
)

func TestClose_fa93199dd8(t *testing.T) {
	t.Run("Test case 1: Normal Close", func(t *testing.T) {
		wc := new(writeCloser)
		err := wc.Close()
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	})

	t.Run("Test case 2: Close when buffer is already nil", func(t *testing.T) {
		wc := new(writeCloser)
		err := wc.Close()
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	})
}
