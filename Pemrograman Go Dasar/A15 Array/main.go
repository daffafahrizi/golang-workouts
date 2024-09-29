package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t \t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//A.15.3. Inisialisasi Nilai Array Dengan Gaya Vertikal
	var fruits2 [4]string

	// cara horizontal
	fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits2 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t \t", len(fruits2))
	fmt.Println("Isi semua element \t", fruits)

	//A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//A.15.5. Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//A.15.6. Perulangan Elemen Array Menggunakan Keyword for
	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruits3); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits3[i])
	}

	fmt.Print("\n")
	//A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range
	var fruits4 = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits4 {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	//A.15.8. Penggunaan Variabel Underscore _ Dalam for - range
	var fruits5 = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits5 {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	//A.15.9. Perulangan Elemen Array Menggunakan Keyword for - range Tanpa Menggunakan Variabel Iterasi
	var fruits6 = make([]string, 2)
	fruits6[0] = "apple"
	fruits6[1] = "manggo"

	fmt.Println(fruits6) // [apple manggo]
}
