package kindling

import "fmt"

// Endpoint is an interface for logging endpoints.
type Endpoint interface {
	Log(message string)
	GetLogLevel() int
	IsStructured() bool
}

type BaseEndpoint struct {
	logLevel     int
	isStructured bool
}

func (b *BaseEndpoint) Log(message string) {
	fmt.Println(message) 
}

func (b *BaseEndpoint) GetLogLevel() int {
	return b.logLevel
}

func (b *BaseEndpoint) IsStructured() bool {
	return b.isStructured
}

func NewEndpoint(logLevel int, isStructured bool) *BaseEndpoint {
	return &BaseEndpoint{
		logLevel:     logLevel,
		isStructured: isStructured,
	}
}
