package main

import "fmt"

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go sequare(squareChan)
	go cube(cubeChan)

	testNum := 3

	fmt.Println("[main] sent testNum to squareChan")
	squareChan <- testNum
	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")
	cubeChan <- testNum
	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal,cubeVal := <-squareChan,<-cubeChan
	sum := squareVal+cubeVal
	fmt.Println("[main] sum of square and cube of",testNum," is" ,sum)
}


func sequare(c chan int){
	fmt.Println("[square] reading")
	num := <-c
	c <- num*num
}


func cube(c chan int){
	fmt.Println("[cube] reading")
	num := <-c
	c <- num*num*num
}