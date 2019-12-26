package protocol

const (
	magicHeaderValue = 0x52654c6f // ReLo written in Hex
)

// Command is an enum type for marking content of packet
type Command int

const (
	// Undefined is a default, unset value
	Undefined Command = iota
	// Register is a client registration request
	Register
	// Confirm is a positive server response for Register request
	Confirm
	// Reject is a negative server response for Register request
	Reject
	// Unregister is an information from client about end of logging provess (possibly application closing)
	Unregister
	// WriteLog is a type of packet containing message log to be saved
	WriteLog
)

// Message represents object sent between client and server
type Message struct {
	Type       Command
	Identifier uint32
	Content    [1024]byte
}
