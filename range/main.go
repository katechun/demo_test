package main

import "fmt"

func main(){
	s:="hello 世界!"
	for _,v:=range s{
		fmt.Println(string(v))
	}
}
