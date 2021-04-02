package main

import "fmt"

type Dog struct {
	Name string
	Call string
}

type Cat struct {
	Name string
	Call string
}

type Animalable interface {
	ToCall()
}

func (d Dog)ToCall(){
	fmt.Printf("%s %s\n",d.Name,d.Call)
}

func (c Cat)ToCall(){
	fmt.Printf("%s %s\n",c.Name,c.Call)
}

func TheToCall(a Animalable){
	a.ToCall()
}

func main(){
	var _ Animalable = (*Dog)(nil)
	var _ Animalable = (*Cat)(nil)
	dog := Dog{"狗蛋","汪汪"}
	cat := Cat{"小花","喵喵"}

	TheToCall(dog)
	TheToCall(cat)
}
