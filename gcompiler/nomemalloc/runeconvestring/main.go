package main

import (
	"fmt"
	"strings"
	"testing"
)

var gogo =strings.Repeat("Go",100)

func f(){_=len([]rune(gogo))} //gogo没有进行拷贝
func g(){_=len([]byte(gogo))}


func main(){
	fmt.Println(testing.AllocsPerRun(1,f))
	fmt.Println(testing.AllocsPerRun(1,g))

}