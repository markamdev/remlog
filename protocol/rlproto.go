package protocol

import (
	"errors"
)

// private const values
const (
	magicHeaderValue     = 0x52654c6f  // ReLo written in Hex, value used when creating network packet
	messageHeaderLen     = (2 + 4 + 2) // 2B for type, 4B for identifier, 2B for content length
	messageContentMaxLen = 1400        // just to be sure that 1500 (typical MTU) will not be exceeded
)

// public (exported) const values
const (
	IdentifierDebug  = (uint32)(0x00000000)
	IdentifierGlobal = (uint32)(0xffffffff)
)

// Command is an enum type for marking content of packet
type Command uint16

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
	Content    []byte
}

// MessageToData converts log message into network packet content (ex. to be sent through UDP)
//
// All values (uintXX) are stored in BigEndian order
func MessageToData(m *Message) ([]byte, error) {
	if m == nil {
		return []byte{}, errors.New("Nil pointer to message given")
	}

	contentLen := len(m.Content)
	// TODO uncomment check below when message content used in all messages (ex. registration)
	/*
		if contentLen == 0 {
			return []byte{}, errors.New("Empty message received")
		}
	*/

	if contentLen > messageContentMaxLen {
		return []byte{}, errors.New("Message content too long")
	}
	result := make([]byte, contentLen+messageHeaderLen)
	// copy type
	result[0] = byte((m.Type >> 8) & 0x00ff)
	result[1] = byte(m.Type & 0x00ff)
	// copy identifier
	result[2] = byte((m.Identifier >> 24) & 0x00ff)
	result[3] = byte((m.Identifier >> 16) & 0x00ff)
	result[4] = byte((m.Identifier >> 8) & 0x00ff)
	result[5] = byte(m.Identifier & 0x00ff)
	// copy message content length (only 2 bytes needed)
	result[6] = byte((contentLen >> 8) & 0x00ff)
	result[7] = byte(contentLen & 0x00ff)
	// copy message content
	for i := 0; i < contentLen; i++ {
		result[8+i] = m.Content[i]
	}

	return result, nil
}

// DataToMessage read bytes received in network packet and creates log message
func DataToMessage(data []byte) (Message, error) {
	if len(data) < messageHeaderLen {
		return Message{}, errors.New("Data buffer too short")
	}
	result := Message{}
	// set type
	result.Type = Command(uint16(data[0])<<8 | uint16(data[1]))
	// set identifier
	result.Identifier = uint32(data[2])<<24 | uint32(data[3])<<16 | uint32(data[4])<<8 | uint32(data[5])
	// copy content
	cLen := len(data) - messageHeaderLen
	if cLen > 0 {
		result.Content = make([]byte, cLen)
		copy(result.Content, data[messageHeaderLen:])
	}
	return result, nil
}
