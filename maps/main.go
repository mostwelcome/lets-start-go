package main

import (
	"fmt"
	"log"
)

type User struct {
	firstName string
	lastName  string
}

func main() {

	//colors : make(map[string]string)

	colors := map[string]string{
		"name":    "Swagata",
		"surname": "Dutta",
	}

	printMap(colors)

	myMap := make(map[string]User)

	me := User{
		firstName: "Swag",
		lastName:  "Dutta",
	}

	myMap["me"] = me

	log.Println(myMap["me"])

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, " is ", hex)
	}
}

// func main() {
//  m := map[string]string{
//    "dog": "bark",
//    "cat": "purr",
//  }

//  for key, value := range m {
//    fmt.Println(value)
//  }
// }
