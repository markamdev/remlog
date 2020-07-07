package main

import (
	"log"

	"github.com/markamdev/remlog/common"
	remlog "github.com/markamdev/remlog/pkg/server"
)

func main() {
	log.Println("RemLog reference server")
	log.Println("Package version:", common.Version)

	cfg := remlog.Config{}
	cfg.AuthPort = common.DefaultAuthPort
	cfg.LogPort = common.DefaultLogPort
	// TODO set some dedicated output function
	// cfg.Ouput = ...
	cfg.DebugMode = true

	err := remlog.InitServer(cfg)
	if err != nil {
		log.Fatal("RemLog refrence server failure: ", err.Error())
	}

	log.Println("RemLog reference server stopped")
}
