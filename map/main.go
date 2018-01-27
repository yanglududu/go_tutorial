package main

import "fmt"

func main() {
	//colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff000",
		"green": "#4bf75",
	}
	colors["white"] = "#fffff"
	delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}