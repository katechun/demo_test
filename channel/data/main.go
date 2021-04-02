package main

import "fmt"

func main(){
	fmt.Println("main() started")
	cc := make(chan chan string)

	go greeter(cc)
	c:=<-cc
	go greet(c)
	c<-"John"

	fmt.Println("main() stopped")
}

func greeter(cc chan chan string){
	c := make(chan string)
	cc <-c
}

func greet(c chan string){
	fmt.Println("Hello "+<-c+"!")
}
