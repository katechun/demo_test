package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var startTime = time.Now()
func main(){

	ctime:=time.Now().Add(3*time.Second)
	ctx,cannel:=context.WithDeadline(context.Background(),ctime)
	defer cannel()
	go seque(ctx,2)
	go seque(ctx,6)
	go seque(ctx,8)

	time.Sleep(5*time.Second)
	fmt.Println("Number of active goroutinmes",runtime.NumGoroutine())




}

func seque(ctx context.Context,seconds int){

	select {
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n",time.Since(startTime).Seconds(),seconds)
		return
		case <-time.After(time.Duration(seconds)*time.Second):
			fmt.Printf("%0.2fs - worker(%ds) completed the job.\n",time.Since(startTime).Seconds(),seconds )
	}

}
