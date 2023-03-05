//package declaration
package main

import "fmt"

/*
package == project == workspace == common source code file
it can contain several file. Condition: every file has to be in the same package
main: Executable() and Reusable

What does package main means?
package main: executable package --> defines a package that can be compiled and then executed
must have a func called main

package blahblah: resuable package --> Defines a package that can be used as a dependency(helper code)
*/
// Go is not OOP
//Import oter packages that we need

//Declare functions, tell go to do things
func main() {
	//fmt.Println("Hi there")

	//var card string = "Ace of Spades"
	//card := "Ace of spades"
	// card := newCard()
	// fmt.Println(card)
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of spades")
	//fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//cards := newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("------------------------------")
	// remainingCards.print()

	// fmt.Println(cards.saveToFile("myCards"))
	// cards := newDeckFromFile("myCards")
	// cards.print()
	cards := newDeck()
	//cards.print()
	fmt.Println("----------------------------------")
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
