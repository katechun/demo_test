package main

import (
	"fmt"
	"testing"
)

var x = []byte{1023:'x'}
var y = []byte{1023:'y'}
var b bool

func f() {b=string(x)!=string(y)}  //不拷贝x
func g(){sx,sy:=string(x),string(y);b=sx==sy}

func main(){
	fmt.Println(testing.AllocsPerRun(1,f))
	fmt.Println(testing.AllocsPerRun(1,g))

}
