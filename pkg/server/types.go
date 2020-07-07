package server

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

// RemLogger defines API used by server engine to save received log
type RemLogger = func(source string, level Severity, message string)

// Config ...
type Config struct {
	// AuthPort is a TCP port number used by clients to register it's id and receive a remlog token
	AuthPort int
	// LogPort is a UDP port number used for receiving log messages from clients
	LogPort int
	// DebugMode flag switches server to verbose mode that prints a lot of messages to standard output
	DebugMode bool
	// Output is an interface to with functions needed to save received messages
	Output RemLogger
}
