package main

import (
"fmt"
"testing"
)

var x = []byte{1023:'x'}
var y = []byte{1023:'y'}
var s string

func f() {s = ("-"+string(x)+string(y))[1:]}  //不拷贝x
func g(){s = string(x)+string(y)}

func main(){
	fmt.Println(testing.AllocsPerRun(1,f))
	fmt.Println(testing.AllocsPerRun(1,g))

}

