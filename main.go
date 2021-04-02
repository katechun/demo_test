package main

import (
	"fmt"
	"log"
)

func main(){
	_,x:=f(1)
	log.Println(x)
	a,_:=echo()
	fmt.Println(a)

}

func f(x int)(_ int,__ int){
	_,__ = x,x
	return
}

func echo()(int,error){
	return 1,nil
}
