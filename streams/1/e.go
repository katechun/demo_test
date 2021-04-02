package main

import (
	"fmt"
	"strings"
	"io"
)


func main(){
	mainSrc := strings.NewReader("Hello Amazing World!")
	src := io.LimitReader(mainSrc,10)
	p:=make([]byte,3)

	for {
		n,err := src.Read(p);
		fmt.Printf("%d bytes read, data: %s\n",n,p[:n])

		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break;
		}else if err != nil {
			fmt.Println("Oops! Some error occured!",err)
			break;
		}
	}
}
