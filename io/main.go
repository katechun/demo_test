package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	var mem runtime.MemStats
	log.Println("memory baseline...")
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	log.Println("memory compatison...")
	runtime.ReadMemStats(&mem)

	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	fmt.Println("go main")

	var sum int
	go func(){

		fmt.Println("goroute")
		for i:=0;i<10000;i++{
			sum=sum+i
		}
		runtime.ReadMemStats(&mem)

		fmt.Println(mem.Alloc)
		fmt.Println(mem.TotalAlloc)
		fmt.Println(mem.HeapAlloc)
		fmt.Println(mem.HeapSys)
		wg.Done()

	}()

	wg.Wait()

}
