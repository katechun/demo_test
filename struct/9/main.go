package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		for {
			if runtime.NumGoroutine() > 3 {
				time.Sleep(20 * time.Microsecond)
			}
			break
		}

		wg.Add(1)

		go func(i int) {

			defer wg.Done()
			fmt.Printf("i = %d\n", i)
		}(i)

		wg.Wait()
	}
}
