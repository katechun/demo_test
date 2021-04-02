package main

import (
	"log"
	"fmt"
	"net/http"
	"path/filepath"
)

var tmpDir = filepath.FromSlash("/root/tmp/")

func main(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request) {
		fmt.Fprint(res,"Hello Golang!")
	})

	http.HandleFunc("/index.html",func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res,req,filepath.Join(tmpDir,"/index.html"));
	})

	log.Fatal(http.ListenAndServe(":9000",nil))
}
