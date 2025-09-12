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
	// We start the first goroutine and increment the WaitGroup counter:
	go printNumbers("Thread 1")

	// We start the second goroutine:
	go printNumbers("Thread 2")

	// We wait for both goroutines to finish:
	wg.Wait()

	// End of the main thread.
	fmt.Println("Main Thread Finished")
}
