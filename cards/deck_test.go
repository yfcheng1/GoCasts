package main

import (
	"os"
	"testing"
)

// func to test the new deck fn
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d)-1])
	}
}

// func test save and pull from file fn

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// delete file in dir
	os.Remove("_decktesting")
	// create a deck and save to file
	d := newDeck()
	d.saveToFile("_decktesting")
	// load from file
	loadedDeck := newDeckFromFile("_decktesting")
	// start compare
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length from file of 52, but got %v", len(loadedDeck))
	}
	// delete file in dir again
	os.Remove("_decktesting")
}
