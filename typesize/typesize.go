package main

import (
	"fmt"
	"unsafe"
)

type T1 struct{
	a struct{}
	x int64
}

type T2 struct{
	x int64
	a struct{}
}

func main(){
	a1 := T1{}
	a2 := T2{}
	fmt.Printf("zero size struct{} of T1 size: %d; T2 (as final field) size: %d",unsafe.Sizeof(a1),unsafe.Sizeof(a2))
}

type tooMuchPadding struct {
	i16 int16
	i64 int64
	i8 int8
	i32 int32
	ptr *string
	b bool
}