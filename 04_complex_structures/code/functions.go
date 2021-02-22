package main

import "fmt"

func printAge(age int) (ageOfSally int, ageOfBob int) {
	ageOfSally = age * 2
	ageOfBob = age * 3
	return
}

// Can use spread operator to show undefind number of args
func printAges(ages ...int) {
	fmt.Println(ages)
}

func main() {
	fmt.Println(printAge(8))
	printAges(5, 10, 15, 20)
}
