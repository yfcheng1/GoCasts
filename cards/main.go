package main

// import "fmt"

func main() {

	cards := newDeck()
	// cards.print()

	hand, remainingCard := deal(cards, 5)
	hand.print()
	remainingCard.print()
}

