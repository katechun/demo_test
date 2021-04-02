package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string){
	for i:=0; i<9000000;i++{
		time.Sleep(100*time.Microsecond)
		fmt.Println(s)
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxx",runtime.NumGoroutine())
		fmt.Println("cccccccccccc",runtime.NumCPU())
	}
}
func main(){
	for i:=0;i<100000;i++{
		go say("katechun")

	}

	say("xc")
}
