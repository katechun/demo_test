package main

import (
	"fmt"
	"io"
	"strings"
)
var Reader io.Reader
func CReader() io.Reader {
	return  Reader
}

func main(){
//	rand := CReader()
	s:=strings.NewReader("aa")
	seed := make([]byte,32)
	_,err := io.ReadFull(s,seed)
	if err != io.EOF {
		panic(err)
	}
	fmt.Println(seed)
}
