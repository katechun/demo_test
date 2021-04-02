package main

import "fmt"

type People struct {
	Name string
}

func (p *People)String()string{
	s:=fmt.Sprintf("print: %v",p)
	return s
}

func main(){
	p := &People{}
	p.String()

}
