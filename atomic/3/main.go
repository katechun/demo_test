package main

import (
	"sync/atomic"
	"sync"
	"fmt"
)

type Config struct {
	a []string
	b map[string]string
}


func main(){
	var v atomic.Value

	go func(){
		var i int
		var a string = "m"
		for {
			i++
			cfg := &Config{
				a : []string{a,a+"1",a+"2",a+"3",a+"4",a+"5"},
			}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n:=0;n<4;n++{
		wg.Add(1)
		go func(){
			for n:=0;n<100;n++{
				cfg := v.Load()
				fmt.Println(cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
