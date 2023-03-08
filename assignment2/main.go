package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLenght float64
}

func printArea(s shape) {
	fmt.Println("Ares: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func main() {
	t := triangle{height: 4, base: 5}
	s := square{sideLenght: 4}

	printArea(t)
	printArea(s)

}
