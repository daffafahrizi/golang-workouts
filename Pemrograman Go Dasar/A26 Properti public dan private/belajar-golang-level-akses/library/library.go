package library
// import "fmt"

// func SayHello(name string) { //public access karena functiionnya diawali huruf besar
// 	fmt.Println("hello")
// 	introduce(name)
// }

// func introduce(name string) { // private access karena functionnya diawali huruf kecil
// 	fmt.Println("nama saya", name)
// }

type Student struct {
	Name string
	Grade int
}