package main

import (
	"fmt"
	"unsafe"
)

type str struct {
	data string
}
func main(){
	//s:=str{}
	//strptr:=unsafe.Pointer(&s)
	//uiptr:=uintptr(strptr)
	//fmt.Println(strptr,uiptr)
	//fmt.Println(unsafe.Pointer(uiptr))

	f:=float64(1.55)
	u:=Float64bites(f)
	fu:=Float64frombits(u)
	fmt.Println(fu)
	s:=ByteSlice2String([]byte("a"))
	fmt.Println(s)
}


func Float64bites(f float64)uint64{
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64frombits(b uint64)float64{
	return *(*float64)(unsafe.Pointer(&b))
}

func ByteSlice2String(bs []byte)string{
	return *(*string)(unsafe.Pointer(&bs))
}