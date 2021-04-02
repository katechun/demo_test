package main

import "fmt"

func main(){
	c:=make(chan int ,3)

	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n",len(c),cap(c))

	for val := range c {
		fmt.Printf("Length of channel c afer value '%v' read is %v\n",val,len(c))
	}
}

func sender(c chan int){
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	close(c)

}
