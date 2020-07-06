package client

// Severity ...
type Severity int

const (
	// Verbose ...
	Verbose Severity = iota
	// Debug ...
	Debug
	// Info ...
	Info
	// Warning ...
	Warning
	// Error ...
	Error
	// Fatal ...
	Fatal
)

// Config ...
type Config struct {
	AuthServer string
	LogServer  string
	LogLevel   Severity
	DebugMode  bool
}
