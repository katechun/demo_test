package main

import (
	"fmt"
)

func main(){

	fmt.Println("main() started")
	c:= make(chan int,3)
	go squares(c)

	c<-1
	c<-2
	c<-3
	c<-4
	c<-1
	c<-2
	c<-3
	c<-4
	c<-1
	c<-2
	c<-3
	c<-4
	c<-1
	c<-2
	c<-3
	c<-4
	c<-1
	c<-2
	c<-3
	c<-4
	fmt.Println("main() stoped")


}

func squares(c chan int){
	for v:=range c {
		fmt.Println(v)
	}
}