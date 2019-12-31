package main

import (
	"flag"
	"fmt"

	remlog "github.com/markamdev/remlog/client"
)

const (
	defServerAddress = "localhost:9999"
	defClientName    = "TesterApp"
)

func main() {
	var serverAddress, clientName string

	// Read command line params (or use default ones)
	flag.StringVar(&serverAddress, "s", defServerAddress, "RemLog server in format 'address:port'")
	flag.StringVar(&clientName, "n", defClientName, "Client name/identifier sent in registration request")
	flag.Parse()

	cnf := remlog.RLCconfig{Server: serverAddress, Name: clientName}
	if err := remlog.Init(&cnf); err != nil {
		fmt.Println("Failed to initialize client: ", err.Error())
		return
	}

	if err := remlog.Register(); err != nil {
		fmt.Println("Failed to register client: ", err.Error())
		return
	}

	fmt.Println("Client successfully registered")

	// try to send messages with different logging levels
	remlog.LogVerbose("This is a Verbose message")
	remlog.LogDebug("This is a Debug message")
	remlog.LogInfo("This is a Info message")
	remlog.LogError("This is a Error message")
	remlog.LogFatal("This is a Fatal message")
}
