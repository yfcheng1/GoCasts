package main

func main() {

	cards := newDeck()
	// cards.print()

	// hand, remainingCard := deal(cards, 5)
	// hand.print()
	// remainingCard.print()

	// cards.toString()
	// fmt.Println(cards.toString())

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting)) // [72 105 32 116 104 101 114 101 33]

	cards.saveToFile("my_cards")
}
