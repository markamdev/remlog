package main

import (
	"fmt"

	"github.com/markamdev/remlog/common"
	"github.com/markamdev/remlog/pkg/client"
)

func main() {
	fmt.Println("RemLog sample client application. Version:", common.Version)
	err := client.InitClient(client.Config{})
	if err != nil {
		panic(err)
	}
}
