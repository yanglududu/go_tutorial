package main

import (
	"fmt"
)

func main() {
	card := newDeck()
	card.shuffle()
	fmt.Println(card.toString())
	card.saveToFile("my_card")
}