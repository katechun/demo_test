package main

import (
	"fmt"
)

func main(){
	fmt.Println("main() started")
	var chan1 chan string

	go service(chan1)

	select{
	case res := <-chan1:
		fmt.Println("Response from service of chan1",res)
	//default:
	//	fmt.Println("No Response")
	}
	fmt.Println("main() stopped")
}

func service(c chan string){
	c<-"response"
}
