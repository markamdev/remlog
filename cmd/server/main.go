package main

import (
	"flag"
	"fmt"

	rls "github.com/markamdev/remlog/server"
)

const (
	defLogFile = "default.log"
	defUDPPort = 9999
)

func main() {
	var serverPort uint
	var outputPath string
	printHelp := false

	flag.UintVar(&serverPort, "p", defUDPPort, "Server listening port")
	flag.StringVar(&outputPath, "o", defLogFile, "Output file for collected logs")
	flag.BoolVar(&printHelp, "h", false, "Print help message")
	flag.Parse()

	if printHelp {
		flag.PrintDefaults()
		return
	}

	// prepare configuration
	cnf := rls.RLSconfig{Port: serverPort, OutFile: outputPath}

	// initialize server library/package
	err := rls.Init(&cnf)
	if err != nil {
		fmt.Println("Failed to init RemLog server:", err.Error())
		return
	}

	// start RemLog server listener
	err = rls.Start(true)
	if err != nil {
		fmt.Println("Failed to start RemLog server:", err.Error())
		return
	}
}
