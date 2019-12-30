package main

import (
	"fmt"

	remlog "github.com/markamdev/remlog/client"
)

func main() {
	cnf := remlog.RLCconfig{Server: "localhost:9999"}
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
