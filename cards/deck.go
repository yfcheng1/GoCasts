package main

import "fmt"

// create a new type of 'deck', which is a slice of strings (like a class)
type deck []string

// receiver type deck (d deck)
// d: actual copy of the deck
// deck: every variable inside deck get access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
