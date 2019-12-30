package client

import (
	"errors"
	"fmt"
	"net"
	"time"

	rlp "github.com/markamdev/remlog/protocol"
)

// RLCconfig structure stores RemLog client configuration (ex. server address and port)
type RLCconfig struct {
	Server string
}

type rlccontext struct {
	valid      bool
	link       *net.UDPConn
	identifier uint32
}

var clientContext rlccontext

// Severity is an enum type for setting log message severity level
type Severity int

const (
	// Verbose is a most intense logging level
	Verbose Severity = iota
	// Debug is a debugging message, should not be used/visible in production environment
	Debug
	// Info is and informative message, used for application tracking in production environment
	Info
	// Error is a message about recoverable application error
	Error
	// Fatal is a message about critical application error, probably the last before application closure
	Fatal
)

// Init function initializes RemLog client
func Init(cnf *RLCconfig) error {
	if cnf == nil {
		return errors.New("Nil config given to client Init() function")
	}
	fmt.Println("Setting up RemLog server as:", cnf.Server)

	// prepare UDP 'connection'
	realAddr, err := net.ResolveUDPAddr("udp", cnf.Server)
	if err != nil {
		return errors.New("Failed to resolve given server address")
	}
	conn, err := net.DialUDP("udp", nil, realAddr)
	if err != nil {
		return errors.New("Failed to create UDP binding")
	}

	// set proper client context field
	clientContext.valid = true
	clientContext.link = conn

	return nil
}

// Register tries to perform 'handshake' with log server
func Register() error {
	if !clientContext.valid {
		return errors.New("RemLog client not initialized")
	}

	msg := rlp.Message{}

	msg.Identifier = rlp.IdentifierGlobal
	msg.Type = rlp.Register
	// TODO remove string copying below as message content for registration not needed
	msg.Content = []byte("Welcome message")

	data, err := rlp.MessageToData(&msg)
	if err != nil {
		return errors.New("Failed to prepare registration message")
	}

	// TODO Registration should be enclosed in some loop as UDP transport is used
	n, err := clientContext.link.Write(data)
	if err != nil || n != len(data) {
		return errors.New("Failed to send registration message")
	}

	// wait for confirmation
	buffer := make([]byte, 1024)
	clientContext.link.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err = clientContext.link.Read(buffer)
	if err != nil || n == 0 {
		return errors.New("Failed to get server response - client not registered")
	}
	resp, err := rlp.DataToMessage(buffer[:n])
	if err != nil {
		return errors.New("Failed to parse server response - client registration not completed")
	}
	// server should response with proper content and client id in message
	if resp.Type == rlp.Confirm {
		clientContext.identifier = resp.Identifier
	} else if resp.Type == rlp.Reject {
		return errors.New("Client registration rejected by server")
	} else {
		return errors.New("Invalid server response - client not registered")
	}

	return nil
}

// SendLog sends log with given severity to configured server, return error if any occured
func SendLog(lvl Severity, message string) error {
	return errors.New("Client.SendLog() Not implemented")
}
