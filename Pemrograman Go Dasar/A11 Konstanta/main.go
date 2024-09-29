package main

import "fmt"

func main() {
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	//DEKLARASI MULTI KONSTANTA
	const (
		square          = "kotak"
		isToday bool    = true
		numeric uint8   = 1
		floatNum        = 2.2
	)

	fmt.Printf("square: %s, isToday: %t, ", square, isToday)
}
