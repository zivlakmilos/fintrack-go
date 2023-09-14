package main

import (
	"fmt"

	"github.com/zivlakmilos/fintrack-go/pkg/core"
)

func main() {
	config, err := core.LoadConfig()
	if err != nil {
		fmt.Printf("got error %v", err)
		return
	}

	fmt.Printf("%v", config)
}
