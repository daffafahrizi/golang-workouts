package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	lastName2 := "wick2"

	fmt.Printf("halo %s %s!\n", firstName, lastName2)

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName3 = "john3"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName3 := "wick3"
	fmt.Printf("halo %s %s!\n", firstName3, lastName3)

	// lastName := "wick"
	// lastName = "ethan"
	// lastName = "bourne"

	//DEKLARASI MULTI VARIABEL
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	
	//deklarasi menggunakan new untuk pointer
	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""
}
