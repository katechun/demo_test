package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}){
	switch msg:=v.(type){
	case *student,student:
		s:=msg.(student)
		fmt.Println(s.Name)
	}
}


func main(){
	s := student{}
	zhoujielun(s)
}