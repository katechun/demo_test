package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("[main] main() started")
	tasks :=make(chan int,10)
	results := make(chan int,10)

	for i := 0;i<3;i++{
		go aqrWorker(tasks,results,i)
	}

	for i:=0;i<5;i++{
		tasks <- i*2
	}

	fmt.Println("[main] Wroke 5 tasks")
	close(tasks)

	for i:=0;i<5;i++ {
		result:=<-results
		fmt.Println("[main] Result",i,":",result)
	}

	//for result:=range results{
	//	fmt.Println("[main] Result",":",result)
	//}

	fmt.Println("[main] main() stopped")
}


func aqrWorker(tasks <-chan int,results chan<- int,instance int){
	for num := range tasks {
		time.Sleep(time.Microsecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n",instance,num)
		results <-num*num
	}
}