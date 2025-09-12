package main

import (
	"fmt"
	"time"
)

func printNumbers(name string) {
	// defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	printNumbers("Main Thread")

	printNumbers("Main Thread Again")

	fmt.Print("Main Thread Finished")
}
