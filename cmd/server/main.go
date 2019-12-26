package main

import (
	"fmt"

	rls "github.com/markamdev/remlog/server"
)

func main() {
	// prepare configuration
	cnf := rls.RLSconfig{Port: 9999}

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
