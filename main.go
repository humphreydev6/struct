package main

import "fmt"

type contactInfo struct {
	email   string
	foneNum int
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jimmy := person{
		firstName: "Joe",
		lastName:  "Tiller",
		contactInfo: contactInfo{
			email:   "perryose71@gmail.com",
			foneNum: 07067351060,
			zipCode: 005553,
		},
	}

	jimmy.updateName("Mark")
	jimmy.print()
}

func (pointerTOPerson *person) updateName(newFirstName string) {
	(*pointerTOPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
