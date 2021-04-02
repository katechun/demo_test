package main

import (
	"fmt"
	"encoding/json"
)

type Student map[string]interface{}

func main(){
	john := Student{
		"FirstName":"John",
		"lastName":"Doe",
		"Age":21,
		"HeightInMeters":1.75,
		"IsMale":true,
	}

	johnJSON ,_ := json.Marshal(john)
	fmt.Println(string(johnJSON))
}
