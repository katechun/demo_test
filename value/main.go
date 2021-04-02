package main

import (
	"fmt"
)

func A(a interface{}){
	fmt.Println(a)
}

func main(){
	b:=666
	A(b)
	A(&b)
}



