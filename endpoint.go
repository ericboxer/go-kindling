package kindling

// EndpointInterface is an interface for logging endpoints.
type EndpointInterface interface {
	Log(message string)
	GetLogLevel() int
	IsStructured() bool
}

type Endpoint struct {
	logLevel     int
	isStructured bool
	logFunc func(message string)
}

func (e *Endpoint) Log(message string) {
	e.logFunc(message)
}

func (e *Endpoint) GetLogLevel() int {
	return e.logLevel
}

func (e *Endpoint) IsStructured() bool {
	return e.isStructured
}

func NewEndpoint(logLevel int, isStructured bool, logFunc func(message string)) *Endpoint {

	return &Endpoint{
		logLevel:     logLevel,
		isStructured: isStructured,
		logFunc:      logFunc,
	}
}
