package main

import (
	"flag"
	"fmt"
	"time"

	remlog "github.com/markamdev/remlog/client"
)

const (
	defServerAddress = "localhost:9999"
	defClientName    = "TesterApp"
	defLoopCount     = 1
)

func main() {
	var serverAddress, clientName string
	var loopCount uint
	var printHelp bool

	// Read command line params (or use default ones)
	flag.StringVar(&serverAddress, "s", defServerAddress, "RemLog server in format 'address:port'")
	flag.StringVar(&clientName, "n", defClientName, "Client name/identifier sent in registration request")
	flag.UintVar(&loopCount, "l", defLoopCount, "Repeat sending set log messages N times")
	flag.BoolVar(&printHelp, "h", false, "Print help message")
	flag.Parse()

	if printHelp {
		flag.PrintDefaults()
		return
	}

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
	for loopCount > 0 {
		fmt.Println("Sending log messages of all available levels")
		remlog.LogVerbose("This is a Verbose message")
		time.Sleep(200 * time.Millisecond)
		remlog.LogDebug("This is a Debug message")
		time.Sleep(200 * time.Millisecond)
		remlog.LogInfo("This is a Info message")
		time.Sleep(200 * time.Millisecond)
		remlog.LogError("This is a Error message")
		time.Sleep(200 * time.Millisecond)
		remlog.LogFatal("This is a Fatal message")
		time.Sleep(1000 * time.Millisecond)
		loopCount--
	}
}
