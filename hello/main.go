package main

import "log"

type myStruct struct {
	FirstName string
}

// func printFirstName() string {}
// `(m *myStruct)` - is a receiver. It ties function to `myStruct`. 
func (m *myStruct) printFirstName() string {
	return m.FirstName
}


func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar:", myVar.printFirstName())
	log.Println("myVar2", myVar2.printFirstName())
}
