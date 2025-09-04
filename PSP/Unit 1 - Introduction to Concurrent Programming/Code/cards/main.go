package main

import (
	"fmt"
	"time"
)

func main() {
	card := "Ace of Spades"

	for range 5 {
		time.Sleep(3 * time.Second)
		fmt.Println(card)
	}
}
