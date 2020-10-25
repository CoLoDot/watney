package utils

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ENV", "TEST")
	os.Setenv("KEY", "myKey")
}

func TestGetEnvVariable(t *testing.T) {
	actual := GetEnvVariable("KEY")
	expected := "myKey"
	if actual != expected {
		t.Errorf("TestGetEnvVariable was incorrect, got: %s, want: %s.", actual, expected)
	}
}
