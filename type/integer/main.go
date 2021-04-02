package main

import (
	"fmt"
	"reflect"
)

func main(){


	//m:=make(map[string][]string)
	//
	//
	//m["a"]=[]string{"1","2"}
	//m["b"]=[]string{"3"}
	//m["c"]=[]string{"4"}
	//fmt.Println(len(m))
	//a:="日本語"                                 // UTF-8 input text
	//b:=`日本語`                                 // UTF-8 input text as a raw literal
	//c:="\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
	//d:="\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
	//e:="\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
	//
	//fmt.Println(a,b,c,d,e)


	type (
		A0 = []string
		A1 = A0
		A2 = struct{ a, b int }
		A3 = int
		A4 = func(A3, float64) *A0
		A5 = func(x int, _ float64) *[]string
	)

	type (
		B0 A0
		B1 []string
		B2 struct{ a, b int }
		B3 struct{ a, c int }
		B4 func(int, float64) *B0
		B5 func(x int, y float64) *A1
	)

	type	C0 = B0

	m:=make(map[string][]string)
	m["a"]=[]string{"1"}
	m["b"]=[]string{"2"}


	m1:=make(map[string][]string)
	m1["a"]=[]string{"1"}
	m1["b"]=[]string{"2"}

	f:=reflect.DeepEqual(m,m1)
	fmt.Println(f)


}
