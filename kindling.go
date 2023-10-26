package kindling

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// Logger manages structured logging with support for multiple logging endpoints and log levels.
type Logger struct {
	level     int
	Endpoints []EndpointInterface
}

type logMessage struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Level   string `json:"level"`
	Trace   string `json:"trace,omitempty"`
}

// NewLogger creates a new Logger with the specified log level.
func NewLogger(level int) *Logger {
	return &Logger{
		level: level,
	}
}

// RegisterEndpoint adds a new logging endpoint to the Logger.
func (l *Logger) RegisterEndpoint(endpoint EndpointInterface) {
	l.Endpoints = append(l.Endpoints, endpoint)
}

// UnregisterEndpoint removes a logging endpoint from the Logger.
func (l *Logger) UnregisterEndpoint(endpoint EndpointInterface) {
	for i, e := range l.Endpoints {
		if e == endpoint {
			l.Endpoints = append(l.Endpoints[:i], l.Endpoints[i+1:]...)
			return
		}
	}
}

func getTrace() string {
	// Capture the caller of the log methods
	pc2, file2, line2, ok2 := runtime.Caller(2)
	if !ok2 {
		return "unknown source"
	}
	fn2 := runtime.FuncForPC(pc2)

	// Capture the function that called the calling function
	pc3, file3, line3, ok3 := runtime.Caller(3)
	if !ok3 {
		return fmt.Sprintf("%s:%d %s()", file2, line2, fn2.Name())
	}
	fn3 := runtime.FuncForPC(pc3)

	// traceString := fmt.Sprintf("%s:%d %s() ", file3, line3, fn3.Name(), file2, line2, fn2.Name())
	traceString := fmt.Sprintf("%s:%d %s() ", file3, line3, fn3.Name())
	return traceString
}



// log sends the log message to all registered endpoints that support the given log level.
func (l *Logger) log(level int, message string, trace ...bool) {
    includeTrace := false
    if len(trace) > 0 {
        includeTrace = trace[0]
    }
	levelPadWidth := 5
	lmTime := time.Now().Format("2006-01-02 15:04:05")
	lmLevel := LogLevelString(level)

	traceInfo := ""
	if includeTrace {
		traceInfo = getTrace()
	}

	lm := logMessage{
		Time:    lmTime,
		Message: message,
		Level:   lmLevel,
		Trace:   traceInfo,
	}

	lmString, _ := json.Marshal(lm)
	messageString := fmt.Sprintf("[%-*s] %s > %s", levelPadWidth, lmLevel, lmTime, message)
	if includeTrace {
		messageString = fmt.Sprintf("%s (Trace: %s)", messageString, traceInfo)
	}

	for _, endpoint := range l.Endpoints {
		if level >= l.level && level >= endpoint.GetLogLevel() {
			if endpoint.IsStructured() {
				endpoint.Log(string(lmString))
			} else {
				endpoint.Log(messageString)
			}
		}
	}
}

// Debug logs a message with the DEBUG level.
func (l *Logger) Debug(message string, trace ...bool) { l.log(DEBUG, message, trace...) }

// Info logs a message with the INFO level.
func (l *Logger) Info(message string, trace ...bool) { l.log(INFO, message, trace...) }

// Warn logs a message with the WARN level.
func (l *Logger) Warn(message string, trace ...bool) { l.log(WARN, message, trace...) }

// Error logs a message with the ERROR level.
func (l *Logger) Error(message string, trace ...bool) { l.log(ERROR, message, trace...) }