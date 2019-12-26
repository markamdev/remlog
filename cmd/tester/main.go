package main

import (
	remlog "github.com/markamdev/remlog/client"
)

func main() {
	cnf := remlog.RLCconfig{Server: "localhost:9999"}
	remlog.Init(&cnf)
}
