package main

import (
	"fmt"
	"reflect"
)

type Pubkey interface{
	Address(s string) 
	Bytes(s string)
}

func Address(s string){

	fmt.Println("address:",s)
}

func Bytes(s string){
	fmt.Println("bytes:",s)
}

func getinterface()interface{}{
	return  (*Pubkey)(nil) 
}

func main(){
	i:=getinterface()
	rt := reflect.TypeOf(i)
	fmt.Println(rt.Elem())
	fmt.Println(rt.Kind())
}
