package main

import "fmt"

type S struct {
	data string
}

func (s S)Reade()string{
	return s.data
}

func (s *S)Write(str string){
	s.data = str
}

//func (s *S)Write(str string){
//	s.data =str
//}

func main(){
	sVals := map[int]S{1:{"A"}}
	sVals[1].Reade()
	sVals[1].Write("test")

	sPtrs := map[int]*S{1:{"A"}}
	sPtrs[1].Reade()
	sPtrs[1].Write("test")

	fmt.Println(sVals,sPtrs[1])
}
