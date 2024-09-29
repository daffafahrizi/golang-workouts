package main

import "fmt"

func main() {
	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// var newFruits = fruits[0:2]

	// fmt.Println(newFruits) // ["apple", "grape"]

	// /*
	// 	Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0,
	// 	hingga elemen sebelum indeks ke-2. Elemen yang memenuhi kriteria tersebut akan didapat,
	// 	untuk kemudian disimpan pada variabel lain sebagai slice baru. Pada contoh di atas,
	// 	newFruits adalah slice baru yang tercetak dari slice fruits,
	// 	dengan isi 2 elemen, yaitu "apple" dan "grape".
	// */

	// var aFruits = fruits[0:3] // indeks 0 sampai sebelum 3
	// var bFruits = fruits[1:4] // indeks 1 sampai sebelum 4

	// var aaFruits = aFruits[1:2]
	// var baFruits = bFruits[0:1]

	// fmt.Println(fruits)   // [apple grape banana melon]
	// fmt.Println(aFruits)  // [apple grape banana]
	// fmt.Println(bFruits)  // [grape banana melon]
	// fmt.Println(aaFruits) // [grape]
	// fmt.Println(baFruits) // [grape]

	// // Buah "grape" diubah menjadi "pinnaple"
	// baFruits[0] = "pinnaple"

	// fmt.Println(fruits)   // [apple pinnaple banana melon]
	// fmt.Println(aFruits)  // [apple pinnaple banana]
	// fmt.Println(bFruits)  // [pinnaple banana melon]
	// fmt.Println(aaFruits) // [pinnaple]
	// fmt.Println(baFruits) // [pinnaple]

	// //A.16.4. Fungsi len()
	// fmt.Println(len(fruits)) // 4

	// //A.16.5. Fungsi cap()
	// var fruits1 = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(len(fruits1)) // len: 4
	// fmt.Println(cap(fruits1)) // cap: 4

	// var aFruits1 = fruits1[0:3]
	// fmt.Println(aFruits)       // len: 3
	// fmt.Println(len(aFruits1)) // len: 3
	// fmt.Println(cap(aFruits1)) // cap: 4

	// var bFruits1 = fruits1[1:4]
	// fmt.Println(len(bFruits1)) // len: 3
	// fmt.Println(cap(bFruits1)) // cap: 3

	// var fruits = []string{"apple", "grape", "banana"}
	// var cFruits = append(fruits, "papaya")

	// fmt.Println(fruits)  // ["apple", "grape", "banana"]
	// fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	// var fruits = []string{"apple", "grape", "banana"}
	// var bFruits = fruits[0:2]

	// fmt.Println(cap(bFruits)) // 3
	// fmt.Println(len(bFruits)) // 2

	// fmt.Println(fruits)  // ["apple", "grape", "banana"]
	// fmt.Println(bFruits) // ["apple", "grape"]

	// var cFruits = append(bFruits, "papaya")

	// fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	// fmt.Println(bFruits) // ["apple", "grape"]
	// fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	//A.16.7. Fungsi copy()
	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst) // watermelon pinnaple apple
	// fmt.Println(src) // watermelon pinnaple apple orange
	// fmt.Println(n)   // 3

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	//A.16.8. Pengaksesan Elemen Slice Dengan 3 Indeks
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2

}
