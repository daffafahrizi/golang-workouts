package main
import(
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return	
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)

	//A.53.2. Decode JSON Ke map[string]interface{} & interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age :", decodedData["Age"])

	//A.53.3. Decode Array JSON Ke Array Objek
	var jsonStringArray = `[
		{"Name": "John Wick", "Age": 27},
		{"Name": "Ethan Hunt", "Age": 32}
	]`

	var data3 []User

	var err2 = json.Unmarshal([]byte(jsonStringArray), &data3)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	//A.53.4. Encode Objek Ke JSON String
	var object = []User{{"John Wick", 27}, {"Ethan Hunt", 32}}
	var jsonData2, err3 = json.Marshal(object)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var jsonString2 = string(jsonData2)
	fmt.Println(jsonString2)
}