package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of 'deck', which is a slice of strings (like a class)
type deck []string

// create a new deck
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

// turn deck to string and join
func (d deck) toString() string {
	s := []string(d)
	return strings.Join(s, ", ")
}

// save string to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read string from file
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}
