package main

import (
	"net/http"
	"fmt"
	"log"
)


func main(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request) {
		fmt.Fprint(res,"Hello World!")
	})

	log.Fatal(http.ListenAndServeTLS(":9000","localhost.crt","localhost.key",nil))
}
