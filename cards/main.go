package main


func main() {
	card := deck{}
	card = append(card, "abc")
	card = append(card, "bb")

	card.print()
}