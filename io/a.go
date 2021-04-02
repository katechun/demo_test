package main

import (
	"io"
	"fmt"
)

var Reader io.Reader

func CReader() io.Reader {
	return Reader
}

func Read(b []byte) (n int, err error) {
	fmt.Println("rand Read  test!!!!!!")
	return io.ReadFull(Reader, b)
}


func main(){
  fmt.Println(CReader())
}
