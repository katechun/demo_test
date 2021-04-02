package main

import (
	"fmt"
	"sync"
)

var i int
func main(){
	var wg sync.WaitGroup
	var mux sync.Mutex

	for j:=0;j<1000;j++{
		wg.Add(1)
		go worker(&wg,&mux)
	}

	wg.Wait()

	fmt.Println("value of i after 1000 operations is",i)
}

func worker(wg *sync.WaitGroup,mux *sync.Mutex){
	mux.Lock()
	i=i+1
	mux.Unlock()
	wg.Done()
}