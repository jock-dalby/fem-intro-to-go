package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("I am a string")) //string

	var x = 4
	/*
		fmt.Println(reflect.TypeOf(x * 5.5))          // command-line-arguments: constant 5.5 truncated to integer
	*/
	fmt.Println(reflect.TypeOf(float64(x) * 5.5)) // float64

}
