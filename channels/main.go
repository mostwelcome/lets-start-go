package main

import (
	"log"
	"math/rand"
	"time"
)

const numpool = 1000

func RandomNumber(n int) int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)
	return value
}

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numpool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)
}
