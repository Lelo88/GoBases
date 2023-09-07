package main

import "fmt"

// Create a new type of deck which is a slice of strings.
type deck []string

// This works like a constructor
func newDeck() deck{
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for i := 1; i <= 13; i++ {
			card := fmt.Sprintf("%v of %v", i, suit)
			cards = append(cards, card)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}