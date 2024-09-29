package main

import "fmt"

func main() {
	var point = 1

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. nilai anda %d\n", point)
	}

	var point2 = 8840.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//SWITCH
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//CASE UNTUK BANYAK KONDISI

	var point4 = 6

	switch point4 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//KURUNG KURAWAL PADA CASE DAN DEFAULT
	var point5 = 1

	switch point5 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//SWITCH DENGAN GAYA IF ELSE
	var point6 = 6

	switch {
	case point6 == 8:
		fmt.Println("perfect")
	case (point6 < 8) && (point6 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//FALLTHROUGH
	/*
		Seperti yang sudah dijelaskan sebelumnya, bahwa switch pada Go memiliki perbedaan dengan bahasa lain.
		Ketika sebuah case terpenuhi, pengecekan kondisi tidak akan diteruskan ke case-case setelahnya.
		Keyword fallthrough digunakan untuk memaksa proses pengecekan tetap diteruskan ke case selanjutnya dengan
		tanpa menghiraukan nilai kondisinya, efeknya adalah case di pengecekan selanjutnya
		selalu dianggap true (meskipun aslinya bisa saja kondisi tersebut tidak terpenuhi, akan tetap dianggap true).
	*/

	var point7 = 6
	fmt.Println("FALL THROUGH")
	switch {
	case point7 == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//SELEKSI KONDISI BERSARANG
	var point8 = 0

	if point8 > 7 {
		switch point8 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point8 == 5 {
			fmt.Println("not bad")
		} else if point8 == 3{
			fmt.Println("you need to learn more")
		} else {
			fmt.Println("You can do it!")
			if point8 == 0 {
				fmt.Printf("your point are %d, try harder!", point8)
			}
		}
	}

} 
