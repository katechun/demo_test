package main

import (
	"reflect"
	"unsafe"
)

func main(){
//	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
//	println(p)

	p:="abc"
	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&p))
	println(s)
	println(*hdr)
}
