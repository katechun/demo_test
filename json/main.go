package main

import (
	"fmt"
	"encoding/json"
)


type Student struct {
	FirstName,lastName string
	Email string
	Age int
	HeightInMeters float64
	IsMale bool
}


func main() {
	john := Student {
		FirstName: "john",
		lastName: "Doe",
		Age: 21,
		HeightInMeters: 1.75,
		IsMale: true,
	}

	johnJSON,_ := json.Marshal(john)

	fmt.Println(string(johnJSON))
}
