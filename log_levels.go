package kindling

// Log levels constants.
const (
	DEBUG = 100 + iota*100
	INFO
	WARN
	ERROR
)

var logLevelNames = map[int]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

func LogLevelString(logLevel int) string {
	level, ok := logLevelNames[logLevel]
	if ok {
		return level
	}
	return "Unknown"
}
