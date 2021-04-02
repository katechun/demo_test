package main

import "sync"

func fanInBySelect(m1, m2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			//case m := <-m1:
			//	c <- m
			//case m := <-m2:
			//	c <- m
			//
			case c <- <-m1:
			case c <- <-m2:

			}
		}
	}()
	return c
}

var wg sync.WaitGroup
func main(){
	wg.Add(2)

	c1:=make(chan string)
	c2 := make(chan string)


	fanInBySelect(c1,c2)

	go func(){
		c1 <- "a"
		wg.Done()
	}()

	go func(){
		c2 <- "b"
		wg.Done()
	}()

	wg.Wait()
}