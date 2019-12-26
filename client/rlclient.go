package client

import (
	"fmt"
)

// RLCconfig structure stores RemLog client configuration (ex. server address and port)
type RLCconfig struct {
	Server string
}

// Init function initializes RemLog client
func Init(cnf *RLCconfig) {
	fmt.Println("Setting up RemLog server as:", cnf.Server)
}
