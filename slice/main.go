package main

import (
	"fmt"
	"reflect"
)

func main(){
	p:=pepole{age:10,name:"katechun"}
	p1:=pepole{}
	//p:=pepole{}
	//if p == (pepole{}){
	//	fmt.Println(p)
	//	fmt.Println("The pepole is empty")
	//}

	if reflect.DeepEqual(p1,p){
		fmt.Println("The same struct")
	}else{
		fmt.Println("The diff struct")
	}

	yourcode()
}

type pepole struct {
	age int
	name string
}

type pepole1 struct {
	age int
	name string
}

func yourcode(){
	fmt.Println()
}