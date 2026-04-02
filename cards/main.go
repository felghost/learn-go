package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Go can infer the type based on the value assigned
	// := is only used to initialize a variable for the first time. After that, you can use = to change its value.
	// card = "Five of Diamonds"

	// card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards.print()
	cards := newDeck()
	
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

// func newCard() string {
//	return "Five of Diamonds"
// }
