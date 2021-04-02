package main

import (
	"fmt"
	"strings"
	"io"
)

func main(){
	src := strings.NewReader("Hello Amazing World!")

	buf := make([]byte,14)

	bytesRead1,err1 := io.ReadFull(src,buf)

	fmt.Printf("Bytes read: %d, value: %s, err: %v\n",bytesRead1,buf[:bytesRead1],err1)

	bytesRead2,err2 := io.ReadFull(src,buf)

	fmt.Printf("Bytes read: %d, value: %s, err: %v\n",bytesRead2,buf[:bytesRead2],err2)

	 bytesRead3,err3 := io.ReadFull(src,buf)

        fmt.Printf("Bytes read: %d, value: %s, err: %v\n",bytesRead3,buf[:bytesRead3],err3)
}


