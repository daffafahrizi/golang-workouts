// // package main
// // import "fmt"
// // import "strings"

// // func main() {
// // 	var names = []string{"John", "Wick"} //ini adalah slice
// // 	printMessage("halo", names)
// // }

// // func printMessage(message string, arr []string) {
// // 	var nameString = strings.Join(arr, " ")
// // 	fmt.Println(message, nameString)
// // }

// //A.18.2. Fungsi Dengan Return Value / Nilai Balik

// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func main() {
//     var randomValue int

//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)

//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)

//     randomValue = randomWithRange(2, 10)
//     fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
//     var value = randomizer.Int()%(max-min+1) + min
//     return value
// }

package main

import "fmt"

func main() {
    divideNumber(10, 2)
    divideNumber(4, 0)
    divideNumber(8, -4)
}

func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}