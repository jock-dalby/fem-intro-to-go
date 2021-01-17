package main

import "fmt"

var varNotInMain string = "Variable outside of main()"

func main() {
	fmt.Println(varNotInMain) // "Variable outside of main()"

	var name string = "Beyonce" // optional if specify type if assigning initial value
	fmt.Println(name)

	var nameTwo string      // same as let in JS, must allocate type if no initial value
	var someNumber int      // Will be given default value of 0
	fmt.Println(nameTwo)    // ""
	fmt.Println(someNumber) // 0

	nameThree := "Test"    // shothand syntax, defines variable with value and type without boilerplate, most common
	fmt.Println(nameThree) // "Test"
}
