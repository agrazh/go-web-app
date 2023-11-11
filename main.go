package main

import (
	"log"
	"math/rand"
	"time"
)

const numPool = 100

func calculateValue(intChan chan int) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(numPool)

	// or: 
	// s1 := rand.NewSource(time.Now().UnixNano())
    // r1 := rand.New(s1)
	// randomNumber := r1.Intn(numPool)

	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
