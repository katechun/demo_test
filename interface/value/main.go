package main

import "fmt"

func main(){
	var s Shape

	fmt.Println("value of s is",s)
	fmt.Printf("type of s is %T\n",s)
}


type Shape interface{
	Ared() float64
	Perimeter() float64
}