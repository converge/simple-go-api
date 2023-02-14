package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHealthz(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	Healthz(w, r)

	got := w.Result()
	want := http.StatusOK

	if got.StatusCode != want {
		t.Errorf("got %d, want %d", got.StatusCode, want)
	}
}

func TestVersion(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/version", nil)

	Version(w, r)

	result := w.Result()
	gotBody, err := io.ReadAll(result.Body)
	if err != nil {
		t.Error(err)
	}

	want := map[string]string{
		"version": "1.0.0",
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gotBody, wantJSON) {
		t.Errorf("gotBody %s, want %s", string(gotBody), string(wantJSON))
	}
}
