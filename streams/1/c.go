package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main(){
	src := strings.NewReader("Hello Amazing World!")
	data,_:=ioutil.ReadAll(src)

	fmt.Printf("Read data of length %d: %s\n",len(data),data)
}	
