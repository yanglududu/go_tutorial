package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Println(p.firstName)
}

func main() {
	//var alex person
	//alex.firstName = "alex"
	//alex.lastName = "armor"
	//fmt.Printf("%+v", alex)
	//fmt.Println(alex)

	jim := person{
		firstName: "alex",
		lastName: "dudu",
		contactInfo: contactInfo{
			email: "duduyang@dy.com",
			zipCode: 100010,
		},
	}

	fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("jim")
	jim.print()
}
