package main

import "fmt"

func main() {
	var value = (((2+6)%3)*4 - 2) / 3

	fmt.Printf("nilai: %d\n", value)

	var value2 = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value2 == 2)

	fmt.Printf("nilai %d (%t) \n", value2, isEqual)

	var left = false
	var right = true
	
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("leftReverse \t(%t) \n", leftReverse)
}
