package main

import (
	"fmt"
	"runtime"
	"time"
)

var c =make(chan bool)
func Fmt(){
	counter:=0
	for {
		select {
		case <-c:
			return
			default:
				time.Sleep(100*time.Microsecond)
				counter++

		}

	}
}


func main(){
	go Fmt()
	fmt.Println("before",runtime.NumGoroutine())
	c<-true
	fmt.Println("after",runtime.NumGoroutine())
}
