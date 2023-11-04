package main

import "log"

func main() {
	var color string
	color = "Green"

	log.Println("Color is:", color)
	changeUsingPointer(&color)
	log.Println("Color is:", color)
}

func changeUsingPointer(s *string) {
	log.Println("address:", s)
	newValue := "Red"
	*s = newValue
}
