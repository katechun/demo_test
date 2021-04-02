package main

import (
	"reflect"
	"fmt"
//	"unsafe"
)

func resinterf() interface{}{
	return 2
}

func main(){
	v := resinterf()
	v1:=2
	r:=reflect.ValueOf(v)

	fmt.Println(r)
	fmt.Println(v1)
}
