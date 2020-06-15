package main

import (
	"fmt"

	remlog "github.com/markamdev/remlog/pkg"
)

func main() {
	fmt.Println("RemLog server implementation")
	err := remlog.InitClient(remlog.ClientConfig{})
	if err != nil {
		panic(err)
	}
}
