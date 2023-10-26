package kindling

import (
	"fmt"
	"testing"
)

func consoleLog(message string) {
	fmt.Println(message)
}

func TestNewEndpoint(t *testing.T) {
	tests := []struct {
		level        int
		isStructured bool
	}{
		{DEBUG, true},
		{INFO, false},
		{WARN, true},
		{ERROR, false},
	}

	for _, test := range tests {
		e := NewEndpoint(test.level, test.isStructured, consoleLog)
		if e.GetLogLevel() != test.level {
			t.Errorf("Expected log level to be %d, but got %d", test.level, e.GetLogLevel())
		}

		if e.IsStructured() != test.isStructured {
			t.Errorf("Expected isStructured to be %v, but got %v", test.isStructured, e.IsStructured())
		}
	}
}

func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		level int
	}{
		{DEBUG},
		{INFO},
		{WARN},
		{ERROR},
	}

	for _, test := range tests {
		e := Endpoint{logLevel: test.level}
		if e.GetLogLevel() != test.level {
			t.Errorf("Expected log level to be %d, but got %d", test.level, e.GetLogLevel())
		}
	}
}

func TestIsStructured(t *testing.T) {
	tests := []struct {
		isStructured bool
	}{
		{true},
		{false},
	}

	for _, test := range tests {
		e := Endpoint{isStructured: test.isStructured}
		if e.IsStructured() != test.isStructured {
			t.Errorf("Expected isStructured to be %v, but got %v", test.isStructured, e.IsStructured())
		}
	}
}

// Since the Log method of BaseEndpoint is abstract (it doesn't have an implementation),
// we won't be writing a test for it directly. Instead, you'd write tests for specific implementations
// of the Endpoint interface that provide a concrete Log method.
