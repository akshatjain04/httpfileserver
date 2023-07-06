package httpfileserver_test

import (
	"testing"

	httpfileserver "github.com/schollz/httpfileserver"
)

// TestOptionNoCache_f670b5cd5b checks the behavior of the OptionNoCache function
func TestOptionNoCache_f670b5cd5b(t *testing.T) {
	t.Run("Disable cache", func(t *testing.T) {
		fs := &httpfileserver.FileServer{}
		httpfileserver.OptionNoCache(true)(fs)
		if !fs.OptionDisableCache {
			t.Error("Expected cache to be disabled, but it's not")
		}
	})

	t.Run("Enable cache", func(t *testing.T) {
		fs := &httpfileserver.FileServer{}
		httpfileserver.OptionNoCache(false)(fs)
		if fs.OptionDisableCache {
			t.Error("Expected cache to be enabled, but it's not")
		}
	})
}
