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
	remlog.SendLog(remlog.Verbose, "This is a Verbose message")
	remlog.SendLog(remlog.Debug, "This is a Debug message")
	remlog.SendLog(remlog.Info, "This is a Info message")
	remlog.SendLog(remlog.Error, "This is a Error message")
	remlog.SendLog(remlog.Fatal, "This is a Fatal message")
}
