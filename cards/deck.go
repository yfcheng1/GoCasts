package main

import "fmt"

// create a new type of 'deck', which is a slice of strings (like a class)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver type deck (d deck)
// d: receiver, actual copy of the deck => think about it as "this"
// deck: every variable inside deck get access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function deal return two values (type deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
