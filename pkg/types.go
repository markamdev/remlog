package remlog

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

// ClientConfig ...
type ClientConfig struct {
	AuthServer string
	LogServer  string
	LogLevel   Severity
	DebugMode  bool
}

// LogOutput ...
type LogOutput int

const (
	// OutputConsole ...
	OutputConsole LogOutput = iota
	// OutputFile ...
	OutputFile
	// OutputDatabase ...
	OutputDatabase
)

// ServerConfig ...
type ServerConfig struct {
	AuthPort   int
	LogPort    int
	DebugMode  bool
	OutputType LogOutput
}
