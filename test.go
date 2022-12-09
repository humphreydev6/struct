// You can edit this code!
// Click here and start typing.
/*package main

import "fmt"

func main() {
	mySlices := []string{"Hi", "there", "What", "are", "you", "doing", "today"}
	updatesSlice(mySlices)
	fmt.Println(mySlices)
}
func updatesSlice(s []string) {
	s[2] = "how"
}

package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "how", "are", "you?"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
*/
package main

import "fmt"

type contactInfo struct {
	email   string
	foneNum int
	zipCode int
}

type person struct {
	firstName  string
	middleName string
	lastName   string
	contactInfo
}

func main() {
	jimmy := person{
		firstName:  "Joe",
		lastName:   "Tiller",
		middleName: "Josh",
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
	(*pointerTOPerson).middleName = newFirstName
}

func (per person) print() {
	fmt.Printf("%+v", per)
}

/*
package main

import "fmt"

type location struct {
	longitude float64
	latitude  float64
}

func main() {
	newYork := location{
		latitude:  40.73,
		longitude: -73.93,
	}

	newYork.changeLatitude()

	fmt.Println(newYork)
}

func (lo location) changeLatitude() {
	lo.latitude = 41.0
}
*/
