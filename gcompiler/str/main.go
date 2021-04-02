package main

import (
	"fmt"
	"strings"
	"testing"
)

var gogo = strings.Repeat("Go",1024)

func main() {
	fmt.Println(testing.AllocsPerRun(1,f))
	fmt.Println(testing.AllocsPerRun(1,g))

}


func f(){for range []byte(gogo){}} //没有拷贝gogo

func g(){bs:=[]byte(gogo);for range bs{}}
