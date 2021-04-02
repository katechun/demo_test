package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	fmt.Println("main() started")
	for i:=0;i<10;i++{
		wg.Add(1)
		go serve(&wg,i)
	}

	wg.Wait()
	fmt.Println("main() stopped")
}

func serve(wg *sync.WaitGroup,instance int){

	time.Sleep(2*time.Second)
	fmt.Println("serve called on instace",instance)
	wg.Done()
}
