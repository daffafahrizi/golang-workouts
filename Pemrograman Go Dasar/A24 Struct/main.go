// package main

// import "fmt"

// type student struct {
// 	name  string
// 	grade int
// }

// func main() {
// 	// var s1 student
// 	// s1.name = "john wick"
// 	// s1.grade = 2

// 	// fmt.Println("name  :", s1.name)
// 	// fmt.Println("grade :", s1.grade)

// 	//A.24.3. Inisialisasi Object Struct
// 	// var s1 = student{}
// 	// s1.name = "wick"
// 	// s1.grade = 2

// 	// var s2 = student{"ethan", 2}

// 	// var s3 = student{name: "jason"}

// 	// fmt.Println("student 1 :", s1.name)
// 	// fmt.Println("student 2 :", s2.name)
// 	// fmt.Println("student 3 :", s3.name)

// }

// embeded struct
package main

import "fmt"

type person struct {
	name string
	age  int
	hairType string
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2
	s1.hairType = "curly"

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("hair type   :", s1.person.hairType)
	fmt.Println("grade :", s1.grade)
	//A.24.7. Pengisian Nilai Sub-Struct
	


}
