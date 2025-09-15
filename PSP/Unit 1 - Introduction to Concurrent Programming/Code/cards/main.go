package main

import (
	"fmt"

	"example/cards/deck"
)

type Coords [2]int

func (c Coords) printCoords() {
	fmt.Printf("(%d, %d)\n", c[0], c[1])
}

func (x Coords) sum(y Coords) Coords {
	r := Coords{x[0] + y[0], x[1] + y[1]}
	return r
}

func main() {
	deck := deck.NewDeck()

	deck.Print()

	a := Coords{10, 20}

	b := Coords{30, 40}

	a.printCoords()
	b.printCoords()

	c := a.sum(b)

	c.printCoords()
}
