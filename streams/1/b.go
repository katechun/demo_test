package main

import (
	"fmt"
	"io"
	"strings"
)

func main(){
	src := strings.NewReader("Hello Amazing World!")

	p:= make([]byte,4)

	for {
		n,err := src.Read(p);
		fmt.Printf("%d bytes read,data: %s\n",n,p[:n])

		if err == io.EOF{
			fmt.Println("--end-of-file--")
			break;
		}else if err != nil {
			fmt.Println("Oops! Some error occured!",err)
			break;
		}
	}
}
