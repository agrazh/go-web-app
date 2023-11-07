package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// var myMap map[string]string
	// `[string]` - is the index of the map, `int` - value stored
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"])
	log.Println(myMap["Second"])
	// maps are imutable
	// `myMap := make(map[string]string)`
	// - will throw an exception "no new variables on left side of :="

	// Map of structs
	users := make(map[string]User)

	user1 := User{
		FirstName: "Lucinda",
		LastName:  "Tanner",
	}

	users["user1"] = user1
	log.Println(users["user1"].FirstName)

	// Slice
	var mySlice []int
	numbers := []string{"one", "two", "three", "four"}

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)
	log.Println(numbers[0:2])
}
