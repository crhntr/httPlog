package httplog_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crhntr/httplog"
)

func ExampleWrap() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "Hello, world!")
	})

	logMux := httplog.Wrap(mux)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	logMux.ServeHTTP(w, r)
}

func TestWrap(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "Hello, world!")
	})

	logMux := httplog.Wrap(mux)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	logMux.ServeHTTP(w, r)
}
