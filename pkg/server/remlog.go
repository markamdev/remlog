package server

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	verboseRun   bool
	serverOutput RemLogger
	locker       sync.WaitGroup
)

// InitServer ...
func InitServer(conf Config) error {
	if conf.Output == nil {
		serverOutput = func(client string, level Severity, message string) {
			log.Println("(RemLog) [", client, "] Level:", level, "Message:", message)
		}
	} else {
		serverOutput = conf.Output
	}

	verboseRun = conf.DebugMode
	if verboseRun {
		serverOutput("SERVER", Debug, "Initializing RemLog server package ...")
	}

	if conf.AuthPort == 0 || conf.LogPort == 0 || conf.AuthPort == conf.LogPort {
		return errors.New("Invalid ports configuration")
	}

	// set SIGINT handler
	// TODO

	// launch auth Go-routine
	locker.Add(1)
	go authRoutine(conf)

	// launch log Go-routine
	locker.Add(1)
	go logRoutine(conf)

	// can wait till all routines stopped
	locker.Wait()
	return nil
}

func authRoutine(conf Config) {
	// temporary debug code
	time.Sleep(time.Second * 2)
	// ====
	serverOutput("AUTH", Debug, "Authorization routine stopped")
	locker.Done()
}

func logRoutine(conf Config) {
	// temporary debug code
	time.Sleep(time.Second * 4)
	// ====
	serverOutput("LOG", Debug, "Logging routine stopped")
	locker.Done()
}
