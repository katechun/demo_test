package main

import (
	"fmt"
	"encoding/json"
)

type Profile struct {
	Username string
	followers int
	Grades map[string]string
}

type Student struct {
	FirstName,lastName string
	Age int
	Profile Profile
	Languages []string
}

func main(){
	var john Student
	john = Student{
		FirstName:"John",
		lastName: "Doe",
		Age:21,
		Profile: Profile{
			Username: "johndoe91",
			followers: 1975,
			Grades:map[string]string{"Math":"A","Science":"A+"},
		},
		Languages: []string{"English","French"},
	}
	johnJSON,err := json.MarshalIndent(john,""," ")
	fmt.Println(string(johnJSON),err)
}
