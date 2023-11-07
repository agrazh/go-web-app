### Basics

<details>
<summary>Decision structures</summary>

```
CODE
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

[Go extension by "Go Team at Google" - [marketplace.visualstudio.com]](https://marketplace.visualstudio.com/items?itemName=golang.go)<br/>
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