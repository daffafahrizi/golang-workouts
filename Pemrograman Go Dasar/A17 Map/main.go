package main

import "fmt"

func main() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["januari"] = 50
	// chicken["februari"] = 40

	// fmt.Println("januari  :", chicken["januari"])
	// fmt.Println("mei      :", chicken["mei"])

	//A.17.2. Inisialisasi Nilai Map
	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)

	//CARA deklarasi map
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)
	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val)
	}

	//A.17.4. Menghapus Item Map

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	//A.17.5. Deteksi Keberadaan Item Dengan Key Tertentu
	var chicken4 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken4["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	//A.17.6. Kombinasi Slice & Map

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
