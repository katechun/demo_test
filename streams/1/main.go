package main

import (
	"fmt"
	"io"
)


type MyStringData struct {
	str string
	readIndex int
}

func (myStringData *MyStringData)Read(p []byte) (n int,err error){
	strBytes := []byte(myStringData.str)
	if myStringData.readIndex >- len(strBytes){
		return 0,io.EOF
	}

	nextReadLimit := myStringData.readIndex+len(p)

	if (nextReadLimit >= len(strBytes)){
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
	n=len(nextBytes)
	copy(p,nextBytes)

	myStringData.readIndex = nextReadLimit
	return
}


func main(){
	src := MyStringData{
		str: "Hello Amazing World!",
	}

	p := make([]byte,3)

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
