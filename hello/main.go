package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}


	animals1 := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, elem := range animals1 {
		log.Println(i, elem)
	}


	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Flufy"


	for _, elem := range animals2 {
		log.Println(elem)
	}


	firstLine := "Hello World"
	for i, l := range firstLine {
		log.Println(i, l)
	}
}
