package main

import (
	"com/go-web-app/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.HelperType
	myVar.Name = "Some name"
	myVar.Number = 9
}
