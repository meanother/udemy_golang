package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards.txt")
	// cards := newDeckFromFile("cards.txt")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
