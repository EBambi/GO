package main

func main() {
	cards := deck{"Card1", newCard()}
	cards = append(cards, "Card2")

	cards.print()
}

func newCard() string {
	return "SuperCard"
}
