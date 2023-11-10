package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

func PrintInfo(a Animal) {
	fmt.Println(a.Says(), a.NumberOfLegs())
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "Ugh"
}
func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{Name: "Samson", Breed: "German Shephered"}

	PrintInfo(&dog)

	gorilla := Gorilla{Name: "Jock", Color: "black", NumberOfTeeth: 32}

	PrintInfo(&gorilla)
}

// interface allowed us to build a function that acceps
// two differnt types. These types satisfies inteface requirement
