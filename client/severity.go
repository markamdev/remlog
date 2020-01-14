package client

// Severity is an enum type for setting log message severity level
type Severity int

const (
	// LvlVerbose is a most intense logging level
	LvlVerbose Severity = iota
	// LvlDebug is a debugging message, should not be used/visible in production environment
	LvlDebug
	// LvlInfo is and informative message, used for application tracking in production environment
	LvlInfo
	// LvlError is a message about recoverable application error
	LvlError
	// LvlFatal is a message about critical application error, probably the last before application closure
	LvlFatal
)

func (s Severity) String() string {
	switch s {
	case LvlVerbose:
		return "VERBOSE"
	case LvlDebug:
		return "DEBUG"
	case LvlInfo:
		return "INFO"
	case LvlError:
		return "ERROR"
	case LvlFatal:
		return "FATAL"
	default:
		return "----"
	}
}
