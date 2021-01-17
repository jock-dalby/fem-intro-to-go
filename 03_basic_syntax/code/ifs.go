package main

import (
	"errors"
	"fmt"
)

// => Function returns an error of type Error
func someFunction() error {
	return errors.New("some error")
}

func main() {

	var someVar = 9

	if someVar > 10 { // curly braces required, parentheses are not
		fmt.Println(someVar)
	}

	// ****************************

	if someVar > 100 {
		fmt.Println("Greater than 100")
	} else if someVar == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

	// ****************************
	err := someFunction()

	// check the variable before executing code
	if err != nil {
		fmt.Println(err.Error())
	}

	// assign variable and check the variable before executing code
	if err := someFunction(); err != nil {
		// Because err defined within if block, it is only scoped
		// within the if block, so if want to redeclare anywhere else
		// we can.
		fmt.Println(err.Error())
	}
}
