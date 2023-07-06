package httpfileserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"sync"
	"time"
)

type testFileServer struct {
	route                 string
	optionDisableCache    bool
	optionMaxBytesPerFile int64
	dir                   string
	cache                 sync.Map
}

type testFile struct {
	bytes  []byte
	header http.Header
	date   time.Time
}

func (fs *testFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implement your ServeHTTP logic here
}

func TestServeHTTP_ac8ae9a3ad(t *testing.T) {
	fs := &testFileServer{
		route: "/test",
		optionDisableCache: false,
		optionMaxBytesPerFile: 1000000,
		dir: "testDir",
	}

	// Test case 1: Request with gzip encoding
	req1, err := http.NewRequest("GET", "/test/file1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req1.Header.Set("Accept-Encoding", "gzip")
	rr1 := httptest.NewRecorder()
	fs.ServeHTTP(rr1, req1)
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr1.Header().Get("Content-Encoding") != "gzip" {
		t.Errorf("handler returned wrong content encoding: got %v want %v", rr1.Header().Get("Content-Encoding"), "gzip")
	}

	// Test case 2: Request without gzip encoding
	req2, err := http.NewRequest("GET", "/test/file2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := httptest.NewRecorder()
	fs.ServeHTTP(rr2, req2)
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if encoding := rr2.Header().Get("Content-Encoding"); encoding != "" {
		t.Errorf("handler returned wrong content encoding: got %v want %v", encoding, "")
	}

	// Test case 3: Request with cache disabled
	fs.optionDisableCache = true
	req3, err := http.NewRequest("GET", "/test/file3", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr3 := httptest.NewRecorder()
	fs.ServeHTTP(rr3, req3)
	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if _, ok := fs.cache.Load(strings.TrimPrefix(req3.URL.Path, fs.route)); ok {
		t.Error("cache should be empty")
	}
}
