package main

import (
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer func(){
			if err := recover();err != nil {
				println(err.(string))
			}
		}()
		p()
		wg.Done()
	}()
	println("main")
	wg.Wait()
}

func p(){
	panic("error panic")
}
