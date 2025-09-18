package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Lest guess a number.")
	fmt.Print("Tell me the lower limit:")

	scanner.Scan()
	lower, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The lower value will be %d.\n", lower)

	fmt.Print("Tell me the upper limit:")
	scanner.Scan()
	upper, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The upper value will be %d.\n", upper)
}



