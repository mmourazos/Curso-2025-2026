package main

import (
	"fmt"
	"sync"

	//"sync"
	"time"
)

// We will use a WaitGroup to wait for the goroutines to finish.
var wg sync.WaitGroup

func printNumbers(name string) {
	// defer allows the following statement to be executed when the function returns.
	// wg.Done() decrements the WaitGroup counter by one.
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// We start the first goroutine and increment the WaitGroup counter:
	wg.Add(1)
	go printNumbers("Thread 1")

	// We start the second goroutine:
	wg.Add(1)
	go printNumbers("Thread 2")

	// We wait for both goroutines to finish:
	wg.Wait()

	// End of the main thread.
	fmt.Println("Main Thread Finished")
}
