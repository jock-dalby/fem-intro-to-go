package main

import "fmt"

func main() {
	// ****************************
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// ****************************
	x := 1

	// This will behave like a while loop
	for x <= 10 {
		fmt.Println(x)
		x += 1
	}

	// ****************************
	var mySentence = "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", letter)
	}

	/*
		Index: 0 Letter: 84
		Index: 1 Letter: 104
		Index: 2 Letter: 105
		Index: 3 Letter: 115
		Index: 4 Letter: 32
		Index: 5 Letter: 105
		Index: 6 Letter: 115
		Index: 7 Letter: 32
		Index: 8 Letter: 97
		Index: 9 Letter: 32
		Index: 10 Letter: 115
		Index: 11 Letter: 101
		Index: 12 Letter: 110
		Index: 13 Letter: 116
		Index: 14 Letter: 101
		Index: 15 Letter: 110
		Index: 16 Letter: 99
		Index: 17 Letter: 101
	*/

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", string(letter))
	}

	/*
		Index: 0 Letter: T
		Index: 1 Letter: h
		Index: 2 Letter: i
		Index: 3 Letter: s
		Index: 4 Letter:
		Index: 5 Letter: i
		Index: 6 Letter: s
		Index: 7 Letter:
		Index: 8 Letter: a
		Index: 9 Letter:
		Index: 10 Letter: s
		Index: 11 Letter: e
		Index: 12 Letter: n
		Index: 13 Letter: t
		Index: 14 Letter: e
		Index: 15 Letter: n
		Index: 16 Letter: c
		Index: 17 Letter: e
	*/
}
