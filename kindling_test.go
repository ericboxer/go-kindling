package kindling

import (
	"strings"
	"testing"
)

type mockEndpoint struct {
	logLevel     int
	isStructured bool
	lastMessage  string
}

func (m *mockEndpoint) Log(message string) {
	m.lastMessage = message
}

func (m *mockEndpoint) GetLogLevel() int {
	return m.logLevel
}

func (m *mockEndpoint) IsStructured() bool {
	return m.isStructured
}

func (m *mockEndpoint) getLastMessage() string {
	return m.lastMessage
}

func TestLogger(t *testing.T) {
	loggerDebug := NewLogger(DEBUG)

	// Create a mock endpoint to capture log messages
	mock := &mockEndpoint{
		logLevel:     DEBUG,
		isStructured: false,
	}

	loggerDebug.RegisterEndpoint(mock)

	testCases := []struct {
		logFunc          func(string, ...bool)
		level            int
		message          string
		trace            bool
		containsMessage  string
		containsLevel    string
		containsTrace    string
	}{
		{loggerDebug.Debug, DEBUG, "Debug message", false, "> Debug message", "[DEBUG]", ""},
		{loggerDebug.Info, INFO, "Info message", true, "> Info message", "[INFO ]", "kindling_test.go"},
		{loggerDebug.Warn, WARN, "Warn message", true, "> Warn message", "[WARN ]", "kindling_test.go"},
		{loggerDebug.Error, ERROR, "Error message", false, "> Error message", "[ERROR]", ""},
	}

	for _, tc := range testCases {
		tc.logFunc(tc.message, tc.trace)
		t.Log(tc.message)
		if !strings.Contains(mock.lastMessage, tc.containsMessage) {
			t.Errorf("Expected log message to contain: %s, but got: %s", tc.containsMessage, mock.lastMessage)
		}
		if !strings.Contains(mock.lastMessage, tc.containsLevel) {
			t.Errorf("Expected log message to be at level: %s, but got: %s", tc.containsLevel, mock.lastMessage)
		}
		if tc.trace && !strings.Contains(mock.lastMessage, tc.containsTrace) {
			t.Errorf("Expected log message to contain trace: %s, but got: %s", tc.containsTrace, mock.lastMessage)
		}
	}
}
