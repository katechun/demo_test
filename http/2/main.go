package main

import (
	"fmt"
	"net/http"
	"sync"
	"log"
)

func createServer(name string,port int) *http.Server{
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprint(res,"Hello: "+name)
	})

	server := http.Server{
		Addr: fmt.Sprintf(":%v",port),
		Handler: mux,
	}

	return &server
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func(){
		server := createServer("ONE",9000)
		log.Fatal(server.ListenAndServe())
		wg.Done()
	}()

        go func(){
                server := createServer("TWO",9001)
                log.Fatal(server.ListenAndServe())
                wg.Done()
        }()

	wg.Wait()
}
