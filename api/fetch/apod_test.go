package fetch

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ENV", "TEST")
}

func TestApod(t *testing.T) {
	actual := Apod()
	expected := `{ "title": "", "url": "", "description": ""}`
	if actual != expected {
		t.Errorf("TestApod was incorrect, got: %s, want: %s.", actual, expected)
	}
}
