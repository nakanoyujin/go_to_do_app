package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	sut := NewMux()

	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200")
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("cannnot read body")
	}

	want := `{"status":"ok"}`
	if string(got) != want {
		t.Errorf("cannnot get status code")
	}

}
