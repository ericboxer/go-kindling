package kindling

import (
	"testing"
)

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{DEBUG, "DEBUG"},
		{INFO, "INFO"},
		{WARN, "WARN"},
		{ERROR, "ERROR"},
		{999999, "Unknown"},
	}

	for _, test := range tests {
		got := LogLevelString(test.input)
		if got != test.expected {
			t.Errorf("Expected %s but got %s for input %d", test.expected, got, test.input)
		}
	}
}
