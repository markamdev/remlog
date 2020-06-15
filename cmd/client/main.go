package main

import (
	"fmt"

	remlog "github.com/markamdev/remlog/pkg"
)

func main() {
	fmt.Println("RemLog sample client application")
	err := remlog.InitClient(remlog.ClientConfig{})
	if err != nil {
		panic(err)
	}
}
