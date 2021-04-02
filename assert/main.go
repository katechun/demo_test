package main

import (
	"fmt"
	"io"
)

func main(){
	//var x interface{}=7
	////i:=x.(interface{})
	//fmt.Println(reflect.TypeOf(x))





}

func f(y I){
	r:=y.(io.Reader)
	r.Read([]byte("a"))

	fmt.Println(r)
}


type I interface {m()}