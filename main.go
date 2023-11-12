package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

	// JSON -> struct
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
	}

	fmt.Printf("unmarshalled: %v \n", unmarshalled)

	// struct -> JSON
	var mySlice []Person

	m1 := Person{
		FirstName: "Wally",
		LastName:  "West",
		HairColor: "red",
		HasDog:    false,
	}

	m2 := Person{
		FirstName: "Diane",
		LastName:  "Prince",
		HairColor: "black",
		HasDog:    true,
	}

	mySlice = append(mySlice, m1, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		fmt.Println("error marshaling:", err)
	}

	fmt.Println("marshaled: \n", string(newJson))
}
