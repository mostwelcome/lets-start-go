package main

import "log"

func main() {
	myVar := "Swagata"
	log.Println("Before change: ", myVar)
	changeVar(&myVar)
	log.Println("After change: ", myVar)
}

func changeVar(s *string) {
	newVar := "Tua"
	*s = newVar
}
