package main

import "fmt"

func main(){

	fmt.Println("main() started")
	c:=make(chan string)
	go greet(c)
	c<-"John"
	close(c)

	c<-"Mike"
	fmt.Println("main() stoped")
}

func greet(c chan string){
	<-c
	<-c
}
