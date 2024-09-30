// package main

// import "fmt"

// func main() {
// 	// var numberA int = 4
// 	// var numberB *int = &numberA

// 	// fmt.Println("numberA (value)   :", numberA)
// 	// fmt.Println("numberA (address) :", &numberA)

// 	// fmt.Println("numberB (value)   :", *numberB)
// 	// fmt.Println("numberB (address) :", numberB)

// 	//efek perubahan nilai pointer
// 	var numberA int = 4
// 	var numberB *int = &numberA

// 	fmt.Println("numberA (value)   :", numberA)
// 	fmt.Println("numberA (address) :", &numberA)
// 	fmt.Println("numberB (value)   :", *numberB)
// 	fmt.Println("numberB (address) :", numberB)

// 	fmt.Println("")

// 	numberA = 5

// 	fmt.Println("numberA (value)   :", numberA)
// 	fmt.Println("numberA (address) :", &numberA)
// 	fmt.Println("numberB (value)   :", *numberB)
// 	fmt.Println("numberB (address) :", numberB)
// }

package main

import "fmt"

func main() {
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}
