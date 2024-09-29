package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i, "dan kuadratnya", i*i)
	}

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	// var i = 0
	// for i <5{
	// 	fmt.Println("Angka", i, "dan kuadratnya", i*i)
	// 	i++
	// }

	var i = 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

	//A.14.4. Penggunaan Keyword for - range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}
	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}

	// A.14.5. Penggunaan Keyword break & continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
	/*
		output:
		angka 2
		Angka 4
		Angka 6
		Angka 8
	*/

	//A.14.6. Perulangan Bersarang

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
	/*
		0 1 2 3 4
		1 2 3 4
		2 3 4
		3 4
		4
	*/

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	/*
	output:
	matriks [0][0]
	matriks [0][1]
	matriks [0][2]
	matriks [0][3]
	matriks [0][4]
	matriks [1][0]
	matriks [1][1]
	matriks [1][2]
	matriks [1][3]
	matriks [1][4]
	matriks [2][0]
	matriks [2][1]
	matriks [2][2]
	matriks [2][3]
	matriks [2][4]
	*/
}
