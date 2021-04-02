package main

import (
	"fmt"
	"sync"
)

type Product struct {
	id int
	name string
	value int
}

func main(){
	p:=&sync.Pool{
		New:func() interface{} {
			return 0
		},
	}

	a:=p.Get().(int)
	p.Put(Product{1,"shoose",100})
	p.Put(1)
	p.Put(2)
	p.Put(3)
	b:=p.Get().(Product)
	fmt.Println(b.value)
	fmt.Println(a,b)
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
