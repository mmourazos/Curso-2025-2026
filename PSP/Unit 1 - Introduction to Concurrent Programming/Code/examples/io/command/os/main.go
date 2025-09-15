package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice that holds the command-line arguments passed to the program.
	if len(os.Args) != 2 {
		fmt.Print("Usage: go run main.go <command> <arg1>")
	} else {
		fmt.Printf("Hello %s!", os.Args[1])
	}
}
