package main

func main() {
	cards := newCard()
	// hand, remaining := deal(cards, 5)
	// hand.printFunc()
	// remaining.printFunc()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.printFunc()
}
