package main

import "fmt"

func main(){

	c:=make(chan int)
	c1:=make(chan int)
	var sum int


	for i:=0;i<10;i++{
		go op(c,c1)
		sum+=<-c1
		if sum ==9 {
			close(c)
		}
	}


	for v:=range c{
		fmt.Println(v)
	}


}


func op(c chan int,c1 chan int){

	for i:=0;i<10;i++{
		c<-i
	}
	c1<-1

}

