package utils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	os.Setenv("ENV", "TEST")
}

func TestGetResponse(t *testing.T) {
	client := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Hello world"))
			}))

	defer client.Close()

	actual := string(GetResponse(client.URL))
	expected := "Hello world"

	if actual != expected {
		t.Errorf(`"TestGetResponse was incorrect, got: "%s", want: "%s"`, actual, expected)
	}
}
