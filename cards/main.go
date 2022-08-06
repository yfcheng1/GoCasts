package main

import "fmt"

func main() {
	// (var card string = "Ace of Spades") is equal to (card := "Ace of Spades")

	// [] is called slice in Go, not array; slice should to have all elements the same type
	// cards := []string{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	// // loop through the cards slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
