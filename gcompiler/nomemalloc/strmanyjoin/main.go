package main

import (
	"fmt"
	"testing"
)

var x,y,z,w ="Hello","World!","Let's","Go!"
var s string

func f() {s = x+y+z+w} //分配一次内存
func g(){s=x+y;s+=z;s+=w}

func main(){
	fmt.Println(testing.AllocsPerRun(1,f))
	fmt.Println(testing.AllocsPerRun(1,g))
}
