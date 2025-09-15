package deck

import "fmt"

type Card struct {
	value int
	suite int
}

type Deck [52]Card

func (d Deck) Print() {
	for _, card := range d {
		card.Print()
	}
}

func (c Card) Print() {
	fmt.Printf("%s of %s\n", values[c.value], suites[c.suite])
}

var suites = [4]string{
	"Harts", "Clubs", "Diamonds", "Spades",
}

var values = [13]string{
	"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King",
}

func NewDeck() Deck {
	var deck Deck

	i := 0
	for suite := range suites {
		for value := range values {
			deck[i] = Card{value: value, suite: suite}
			i++
		}
	}

	return deck
}
