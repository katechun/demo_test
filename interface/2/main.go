package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}){
	switch msg:= v.(type) {
	case student:
		fmt.Println(msg.Name)
	}
}


func main(){
	s := student{}

	zhoujielun(s)
}