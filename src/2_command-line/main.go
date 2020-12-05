package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args //type inference
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		var message string = "Hello World" //manual type decalration
		fmt.Println(message)
	}
}
