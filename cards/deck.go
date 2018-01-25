package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + "of" + value)
		}
	}
	return cards
}

func (d deck) print() {
	for _, val := range d {
		fmt.Println(val)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[0: handSize], d[handSize:]
}

func (d deck) toString () string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile (fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}