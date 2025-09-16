package main

import (
	"flag"
	"fmt"
)

/*
* This program demonstrates how to pass arguments to a command in Go.
* We will use the flag module to parse command-line arguments.
 */

func main() {
	// We can read a _flag_ using flag.StringVar or flag.String functions.
	userName := flag.String("name", "Anonymous", "User's name.")

	// var userName string
	// flag.StringVar(&userName, "name", "Anonymous", "User's name.")

	flag.Parse()

	// userName is a pointer to a string, so we need to dereference it.

	fmt.Println("Hello", *userName, "!")

	// We need to use pointers with the flag function StringVar.
	// var userName string
	// flag.StringVar(&userName, "name", "Anonymous", "User's name.")
	//
	// userName is a string, so we don't need to dereference it.
	//
	// fmt.Println("Hello", userName, "!")
}
