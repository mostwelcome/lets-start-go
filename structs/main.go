package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	// swagata := person{firstName: "Swagata", lastName: "Dutta"}
	// fmt.Println(swagata)

	// var swagata person
	// swagata.firstName = "Swagata"
	// swagata.lastName = "Dutta"
	// fmt.Println(swagata)

	// fmt.Printf("%+v", swagata)

	swagata := person{
		firstName: "Swagata",
		lastName:  "Dutta",
		contactInfo: contactInfo{
			email:   "swag@gmail.com",
			zipCode: 711205,
		},
	}
	swagata.print()
	// swagataPointer := &swagata
	// swagataPointer.updateName("Tua")
	swagata.updateName("Tua")
	swagata.print()
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
package main

import "fmt"

func main() {
 name := "bill"

 namePointer := &name

 fmt.Println(&namePointer)
 printPointer(namePointer)
}

func printPointer(namePointer *string) {
 fmt.Println(&namePointer)
}

*/
