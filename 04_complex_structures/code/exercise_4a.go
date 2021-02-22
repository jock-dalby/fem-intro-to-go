package main

import "fmt"

func getAverage(num1, num2, num3 float64) float64 {
	sum := num1 + num2 + num3
	return sum / 3
}

func main() {
	fmt.Println(getAverage(2.2, 3.4, 4.6))
}
