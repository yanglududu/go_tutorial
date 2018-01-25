package main

import (
	"fmt"
)

func main() {
	card := newDeck()
	fmt.Println(card.toString())
	card.saveToFile("my_card")
}