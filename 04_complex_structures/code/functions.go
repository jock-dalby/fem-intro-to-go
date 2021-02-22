package main

import "fmt"

func printAge(age int) (ageOfSally int, ageOfBob int) {
	ageOfSally = age * 2
	ageOfBob = age * 3
	return
}

func main() {
	fmt.Println(printAge(8))
}
