package main

import (
	"io"
	"net/http"
	"net/rpc"
	"test/test/http/common"
)

func main(){
	mit := common.NewCollege()

	rpc.Register(mit)
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		io.WriteString(res,"RPC SERVER LIVE!")
	})

	http.ListenAndServe(":9000",nil)
}
