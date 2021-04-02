package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{}
	v = (*int)(nil)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(v == nil)
}
