package main

import (
	"fmt"
	//"sync"
	"time"
)

// var wg sync.WaitGroup

func printNumbers(name string) {
	// defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// wg.Add(1)
	go printNumbers("Thread 1")
	// wg.Add(1)
	go printNumbers("Thread 2")

	// We wait for both goroutines to finish:
	// wg.Wait()

	select {}
}
