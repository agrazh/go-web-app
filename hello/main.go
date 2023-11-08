package main

import "log"

func main() {
	// var isTrue bool
	isTrue := true
	myNum := 101

	// if isTrue == true
	if myNum >= 100 && isTrue {
		log.Println(isTrue)
	} else if myNum < 100 && !isTrue {
		log.Println(isTrue)
	} else {
		log.Println(isTrue)
	}

	// switch statement
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println(myVar)
	case "dog":
		log.Println(myVar)
	default:
		log.Println(myVar)
	}
}
