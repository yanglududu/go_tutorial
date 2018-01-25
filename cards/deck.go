package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, val := range d {
		fmt.Println(val)
	}
}
