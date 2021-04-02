package main

import "fmt"

func main(){
	//c1:=make(chan int)
	//c2:=make(chan int)
	//
	//fmt.Println(reflect.DeepEqual(c1,c2))
	//fmt.Println(c1==c2)

	ch01 := make(chan int)
	ch02 := make(chan int)
	if ch01 == ch02 {
		fmt.Println("ch01 == ch02")
	} else {
		fmt.Println("ch01 != ch02")
	}
}
