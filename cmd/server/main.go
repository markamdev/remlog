package main

import (
	"fmt"

	"github.com/markamdev/remlog/common"
	remlog "github.com/markamdev/remlog/pkg/server"
)

func main() {
	fmt.Println("RemLog server implementation. Version:", common.Version)
	err := remlog.InitServer(remlog.Config{})
	if err != nil {
		panic(err)
	}
}
