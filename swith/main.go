package main

import (
	"fmt"
	"time"
)

func main(){
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println(hour)
		fmt.Println("Good morning!")
		//break
	case hour < 17:
		fmt.Println(hour)
		fmt.Println("Good afternoon!")
		//break
	default:
		fmt.Println(hour)
		fmt.Println("Good evening!")
	}
}
