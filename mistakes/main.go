package main

import (
	"fmt"
)

func main(){
	in := []int{1,2,3}

	var out []*int
	for _,v := range in {
		out =append(out,&v)
	}

	fmt.Println("Values:",*out[0],*out[1],*out[2])
	fmt.Println("Addresses:",out[0],out[1],out[2])
}
