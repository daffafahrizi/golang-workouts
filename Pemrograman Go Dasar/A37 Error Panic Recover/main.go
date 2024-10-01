package main

import (
	"errors"
	"fmt"
	"strings"
)

// func main(){
// 	var input string
// 	fmt.Print("type some number: ")
// 	fmt.Scanln(&input)

// 	var number int
// 	var err error
// 	number, err = strconv.Atoi(input)

// 	if err == nil {
// 		fmt.Println(number, "is number")
// 	} else {
// 		fmt.Println(input, "is not number")
// 		fmt.Println(err.Error())
// 	}
// }

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// func main() {
//     var name string
//     fmt.Print("Type your name: ")
//     fmt.Scanln(&name)

//     if valid, err := validate(name); valid {
//         fmt.Println("halo", name)
//     } else {
//         fmt.Println(err.Error())
//     }
// }

//A.37.3. Penggunaan panic

// func main() {
//     var name string
//     fmt.Print("Type your name: ")
//     fmt.Scanln(&name)

//     if valid, err := validate(name); valid {
//         fmt.Println("halo", name)
//     } else {
//         panic(err.Error())
//         //fmt.Println("end")
//     }
// }

// A.37.4. Penggunaan recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

// func main() {
// 	defer catch()

// 	var name string
// 	fmt.Print("Type your name: ")
// 	fmt.Scanln(&name)

// 	if valid, err := validate(name); valid {
// 		fmt.Println("halo", name)
// 	} else {
// 		panic(err.Error())
// 		//fmt.Println("end")
// 	}
// }

// A.37.5. Pemanfaatan recover pada IIFE

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()

	panic("some error happen")
}
