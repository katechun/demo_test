package main

import (
	"fmt"
	"sync"
)

type data struct {
	mu sync.Mutex
	name string
}

func (d *data)get() string{
	return d.name
}

func (d *data)setName(name string){
	d.name=name
}

func main(){
	d:=data{name:"katechun"}
	fmt.Println(d.get())

	d.setName("test")

	fmt.Println(d.get())
}
