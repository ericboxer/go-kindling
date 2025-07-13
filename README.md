# Kindling Logger ( A Go implimentation)

A flexible, multi-endpoint Go logger with structured and unstructured output
support.

## Features

- Multiple output endpoints (console, file, network, etc.)
- Structured (JSON) and unstructured logging
- Optional call stack tracing
- Configurable log levels
- Thread-safe operations

## Quick Start

```go
logger := kindling.NewLogger(kindling.INFO)
endpoint := kindling.NewEndpoint(kindling.DEBUG, false, func(msg string) {
    fmt.Println(msg)
})
logger.RegisterEndpoint(endpoint)
logger.Info("Hello, world!")
```
