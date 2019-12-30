package server

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	rlp "github.com/markamdev/remlog/protocol"
)

const (
	readBufferLen = 1500 // 1500 it's a typical MTU in TCP/IP networks so packet should not reach this value
)

// RLSconfig structure containing RemLog server configuration
type RLSconfig struct {
	Port int
}

type rlscontext struct {
	valid    bool
	listener *net.UDPConn
	clients  map[string]uint32
	active   bool
}

var serverContext rlscontext

// Init function initializes RemLog server
func Init(conf *RLSconfig) error {
	fmt.Println("Initializing RemLog server on port:", conf.Port)

	// prepare server context
	var ctx rlscontext
	ctx.clients = make(map[string]uint32)

	// open UDP listening port
	addr := net.UDPAddr{Port: conf.Port}
	connection, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		fmt.Println("Failed to open listening socket")
		return err
	}
	ctx.listener = connection

	// open logging file

	// prepare random generator for id generation
	rand.Seed(time.Now().UnixNano())

	// no errorr till now
	ctx.valid = true
	serverContext = ctx
	return nil
}

// Start function starts UDP listener
//
// wait parameter informs whether function should block on listener (wait == true)
// or exit immediately (wait == false)
func Start(wait bool) error {
	if !serverContext.valid {
		return errors.New("Server not initialized - Init() not called or failed")
	}

	serverContext.active = true

	if wait {
		startListener()
	} else {
		go startListener()
	}

	return nil
}

func startListener() {
	//
	buffer := make([]byte, readBufferLen)

	for serverContext.active {
		n, addr, err := serverContext.listener.ReadFrom(buffer)
		if err != nil {
			// TODO: Handle error
			continue
		}
		if n == 0 {
			// zero-length packet means something is wrong
			// TODO: Handle error
			continue
		}
		// temporary debug message
		fmt.Println("Packet received from:", addr.String())

		msg, err := rlp.DataToMessage(buffer[:n])
		if err != nil {
			fmt.Println("Failed to read log message from packet:", err.Error())
			continue
		}
		switch msg.Type {
		case rlp.Register:
			fmt.Println("Registration request received")
			registerClient(addr, msg)
		case rlp.Unregister:
			fmt.Println("Deregistration request recevied")
		case rlp.WriteLog:
			fmt.Println("Packet with log message received")
		default:
			fmt.Println("Unsupported message type - skipping")
		}

	}

}
