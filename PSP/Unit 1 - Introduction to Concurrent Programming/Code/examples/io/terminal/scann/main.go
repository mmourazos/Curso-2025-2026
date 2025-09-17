package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// In this example we will show how to interact with
	// the terminal asking for inputs to the user.

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert your name:")

	// Now we need to use the scanner to look for user input.
	scanner.Scan()
	fmt.Println("Your names is", scanner.Text()) // scanner.Text() will return the line read.

	fmt.Print("Insert the year you were born:")
	scanner.Scan()

	year, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	age := time.Now().Year() - int(year)

	fmt.Print("Your age is:", age)
}
