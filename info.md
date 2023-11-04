### Basics

<details>
<summary>Types and Structs</summary>

```
CODE
```
</details>

<details>
<summary>Pointers</summary>

`&` - get address of the value. `*` - get value by the address.<br/>
`s *string` - decalare pointer to the string value.


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

```
// package declaration is mandatory
package main

import "fmt"

// main() func is mandatory
func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

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
$ cd /my/parent/dir
$ go work init
$ go work use project-one
$ go work use project-two
```

[VS Code - [code.visualstudio.com]](https://code.visualstudio.com/download)<br/>

[Go binanaries - [go.dev]](https://go.dev/dl/)<br/>
\* Terminal > `$ go version`