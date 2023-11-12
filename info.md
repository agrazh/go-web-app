### Basics

<details>
<summary>Title</summary>

```
CODE
```
</details>

<details>
<summary>Writing tests</summary>

`$ go test`<br/>
`$ go test -v`<br/>
`$ go test -cover`<br/>
`$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out`</br>

`main_test.go` :

```
package main

import "testing"

// //
// Manual approach
// //

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}

// //
// More efficient approch
// //

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected  float32
	isErr    bool
}{
	{"valid-data", 100.0, 5.0, 20.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one:", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
```
</details>

<details>
<summary>Reading and writing JSON</summary>

`json.Unmarshal([]byte(myJson), &unmarshalled)` - deserialize<br/>
`newJson = json.Marshal(mySlice)` - serialize

```
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
```
</details>

<details>
<summary>Channels</summary>

Channels are used for goroutine comunication.

`intChan <- randomNumber` - push into the channel<br/>
`num := <-intChan` - pull from the channel<br/>
`go calculateValue(intChan)` - run goroutine<br/>
`defer close(intChan)` - close channel once parent function is done.

```
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
```
</details>

<details>
<summary>Packages</summary>

`$ cd /app-dir/helpers`<br/>
`$ cat help.go` :
```
package helpers

func MyFunc() {...}  // upper-case name

```

<br/>

`$ cd /app-dir`<br/>
`$ go mod init com/my-web-app`<br/>
`$ cat go.mod` :

```
module com/go-web-app

go 1.18
```

<br/>

`$ cat main.go` :

```
package main

import "com/go-web-app/helpers"

func main() {...}
```

</details>

<details>
<summary>Interfaces</summary>

`type Animal interface {}` - interface allowed us to build a function that acceps two differnt types. These types satisfies inteface requirement.

```
package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

// this function will be able to acccept Dog and Gorrila type
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
```
</details>

<details>
<summary>Loops</summary>

`for i := 0; i <= 10; i++ {}`<br/>
`for i, elem := range animals {}`

```
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
```
</details>


<details>
<summary>Decision structures</summary>

`if {} if else {} else {}`<br/>
`switch {case: default: }`

```
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
```
</details>

<details>
<summary>Maps and Slices</summary>

`myMap := make(map[string]int)` : `[string]` - is the index of the map, `int` - value stored at indexes.<br/>
Map is mutable. Mutable types are passed by reference, imutable ones - by value.

`numbers := []string` - slice of strings

`sort` package

```
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
```
</details>

<details>
<summary>Structs with functions</summary>

`func (m *myStruct) printFirstName() string {}`, where `(m *myStruct)` is a receiver - tt ties function to `myStruct`:

```
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
```
</details>

<details>
<summary>Types and Structs</summary>

Struct type.<br/>
_Uppercase names_ are available outside of the package. _Lowercase names_ are available only inside the package.  

```
package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	birthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Ameera",
		LastName:    "Knowles",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, user.birthDate)
}

func whatever()    {} // is only available only whithin current package
func Whatever()    {} // is visible outside of current pacakge
var special string // is availalbe only whithin current package
var Special string // is available outside of current package
```
</details>

<details>
<summary>Pointers</summary>

`&` - get address of the value. `*` - get value by the address.<br/>
`s *string` - decalare pointer to the string value.<br/>
`log` package.

```
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
```
</details>

<details>
<summary>Variables & Functions</summary>

`fmt` package.

```
// package declaration is mandatory
package main

import "fmt"

// main() func is mandatory
func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int
	// var i = 4
	// i := 4

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 4
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingWasSaid := saySomething()
	fmt.Println("The func returned:", whatWasSaid, theOtherThingWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
```
</details>

<details>
<summary>Toggle block example</summary>
<code style="white-space:nowrap;">Hello World, how is it going?</code>
</details>


<br/>


### docs

[Go standart library documentatoin - [pkg.go.dev]](https://pkg.go.dev/std)<br/>
[Search for open source Go packages - [pkg.go.dev]](https://pkg.go.dev/)<br/>
[Javascript docs - [developer.mozilla.org]](https://developer.mozilla.org/en-US/docs/Web/javascript)<br/>
[CSS docs - [developer.mozilla.org]](https://developer.mozilla.org/en-US/docs/Web/CSS)<br/>
[jsdelivr - CDN which hosts a lot of open source JS and CSS resources - [jsdelivr.com]](https://www.jsdelivr.com/)<br/>
[cdnjs - another great CDN - [cdnjs.com]](https://cdnjs.com/)

Go by Example [[gobyexample.com]](https://gobyexample.com/)

<br/>


### "hello" program

`main.go`:
```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
`$ go run main.go`

<br/>


### Installation

\* Go extension by "Go Team at Google" [[marketplace.visualstudio.com]](https://marketplace.visualstudio.com/items?itemName=golang.go)<br/>
\* VS Code > `Ctrl + Shift + p` > `Go: Install/Update Tolls` > select all > `OK` <br/>
\* "Error loading workspace: gopls was not able to find modules in your workspace" [[stackoverflow.com]](https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory) : </br>

```
$ cd /app-dir/module1
$ go mod init app.org/module1
$ cd /app-dir/module2
$ go mod init app.org/module2

$ cd /app-dir
$ go work init
$ go work use module1
$ go work use module2
```

Multi-module workspaces [[go.dev]](https://go.dev/doc/tutorial/workspaces)

---

VS Code [[code.visualstudio.com]](https://code.visualstudio.com/download)<br/>

Terminal > `$ go version`<br/>
Go binanaries [[go.dev]](https://go.dev/dl/)<br/>