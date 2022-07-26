package main

import (
	"fmt"
	// "encoding/json"
)

type Student struct {
	name string `json:Name`
	age int `json:Age`
}

type Student2 struct {
	name string `json:Name`
	age int `json:Age`
	school string `json:School`
}
func main() {
	frequency := map[string]interface{}{
		"4": "mere",
		"10": -1,
		"12": Student{name:"Dan", age:18},
		"19": Student2{name:"Cosmin", age:25, school:"UPB"},
	}
	
	for _, v := range frequency {
		// if value, ok := v.(Student2); ok {
		// 	fmt.Println(value.name, value.age)
		// }
		switch s := v.(type) {
		case Student:
			fmt.Printf("Este student 1: %s, %d\n", s.name, s.age)
		case Student2:
			fmt.Printf("Este student 2: %s, %d, %s\n", s.name, s.age, s.school)
		case string:
			fmt.Printf("Found string: %s\n", s)
		default:
			fmt.Printf("No match for %v\n", v)
		}
	}

	// var stud Student
	
	// err := json.Unmarshal("{
	// 	"Name": "Dan",
	// 	"Age": 25,
	// 	"School": "UPB"
	// }", &stud)




	// var array = []int{1,2,5,3,4,2,7,10,3}
	// for _, v := range array {
		
	// }
	// frequency[4] = 11
	// frequency[10] = -1


	// for x := range frequency {
	// 	fmt.Println(val)
	// data := map[string]string{}
	// data["ana"] = "mere"

	// x, exists := data["ana"]
	// if !exists {

	// 	fmt.Println("nu exista")
	// }
	// fmt.Println(x)

	// if x, boolvar := data["dana"]; boolvar {
	// 	fmt.Println("nu exista")

	// }



	
	
}