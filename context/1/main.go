package main

import (
	"context"
	"fmt"
	"runtime"
)


func sequ(ctx context.Context){
	counter:=1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			counter++
		}
	}

}
func main(){
	ctx,cannel:=context.WithCancel(context.Background())
	go sequ(ctx)
	fmt.Println("before:",runtime.NumGoroutine())
	cannel()
	fmt.Println("after:",runtime.NumGoroutine())

}
