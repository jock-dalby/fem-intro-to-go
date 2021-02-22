package main

import "fmt"

func getAverage(args ...float64) float64 {
	// Define sum var as a float64
	sum := 0.0
	// sum up all args
	for _, value := range args {
		sum = sum + value
	}
	// find num of args. Must be as float as want to do calculation against float and both sides of calc
	// must be same type
	numOfArgsAsInt := len(args)
	numOfArgsAsFloat := float64(numOfArgsAsInt)
	// Find average of sum / num of args
	return sum / numOfArgsAsFloat
}

func main() {
	fmt.Println(getAverage(2.2, 3.4, 4.6))
}
