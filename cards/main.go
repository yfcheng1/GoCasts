package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCard := deal(cards, 5)
	// hand.print()
	// remainingCard.print()

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards := newDeckFromFile("my")
	// cards.print()

}
