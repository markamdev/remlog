package server

import (
	"errors"
	"fmt"
	"io"
	"net"
)

// RLSconfig structure containing RemLog server configuration
type RLSconfig struct {
	Port int
}

type rlscontext struct {
	valid    bool
	listener *net.UDPConn
	clients  map[string]io.Writer
	active   bool
}

var serverContext rlscontext

// Init function initializes RemLog server
func Init(conf *RLSconfig) error {
	fmt.Println("Initializing RemLog server on port:", conf.Port)

	// prepare server context
	var ctx rlscontext

	// open UDP listening port
	addr := net.UDPAddr{Port: conf.Port}
	connection, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		fmt.Println("Failed to open listening socket")
		return err
	}
	ctx.listener = connection

	// open logging file

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

	if wait {
		startListener()
	} else {
		go startListener()
	}

	return nil
}

func startListener() {

}
