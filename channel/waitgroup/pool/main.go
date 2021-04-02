package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	fmt.Println("[main] main() started")
	var wg sync.WaitGroup

	tasks := make(chan int,10)
	results := make(chan int,10)

	for i:=0;i<3;i++{
		wg.Add(1)
		go sqrWorker(&wg,tasks,results,i)
	}

	for i:=0;i<5;i++{
		tasks <-i*2
	}
	fmt.Println("[main] Wote 5 tasks")

	close(tasks)
	wg.Wait()

	for i:=0;i<5;i++{
		result:=<-results
		fmt.Println("[main] Result",i,":",result)
	}

	fmt.Println("[main] main() stopped")
}

func sqrWorker(wg *sync.WaitGroup,tasks <-chan int,results chan<-int,instance int){
	for num := range tasks {
		time.Sleep(time.Microsecond)
		fmt.Printf("[worker %v] Senging result by worker %v\n",instance ,num)
		results <- num*num
	}
	wg.Done()
}
